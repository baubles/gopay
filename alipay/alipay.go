package alipay

import (
	"bytes"
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"hash"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"
)

var Debug = false

type AlipayClient struct {
	ServerUrl       string
	AppId           string
	SignType        string
	Charset         string
	Format          string
	PrivateKey      string
	AlipayPublicKey string
	httpClient      *http.Client
}

func NewAlipayClient(appId, signType, privateKey, alipayPublicKey string) *AlipayClient {
	return &AlipayClient{
		ServerUrl:       "https://openapi.alipay.com/gateway.do",
		AppId:           appId,
		PrivateKey:      privateKey,
		AlipayPublicKey: alipayPublicKey,
		Format:          FORMAT_JSON,
		Charset:         CHARSET_UTF8,
		SignType:        signType,
		httpClient: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

func (a *AlipayClient) Do(req *AlipayRequest) (interface{}, error) {
	params, err := a.params(req)
	if err != nil {
		return nil, err
	}
	if params.Has(SIGN_TYPE) {
		if err := a.signParams(params); err != nil {
			return nil, err
		}
	}

	if params.Get(CHARSET) != CHARSET_UTF8 {
		return nil, errors.New("[alipay] charset only support utf8")
	}
	reqBody := params.Encode()
	resp, err := a.httpClient.Post(a.ServerUrl, "application/x-www-form-urlencoded;charset="+a.Charset, strings.NewReader(reqBody))
	if Debug {
		log.Printf("[gopay.alipay] request %s\t%s\n", a.ServerUrl, reqBody)
	}
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if Debug {
		log.Printf("[gopay.alipay] response %s\n", respBody)
	}
	err = a.parseResponse(bytes.NewReader(respBody), req)
	return req.response, err
}

func (a *AlipayClient) SdkDo(req *AlipayRequest) (string, error) {
	params, err := a.params(req)
	if err != nil {
		return "", err
	}
	if params.Has(SIGN_TYPE) {
		if err := a.signParams(params); err != nil {
			return "", err
		}
	}

	if params.Get(CHARSET) != CHARSET_UTF8 {
		return "", errors.New("[alipay] charset only support utf8")
	}

	return params.Encode(), nil
}

func (a *AlipayClient) params(req *AlipayRequest) (AlipayParams, error) {
	params := AlipayParams{}
	params.Put(METHOD, req.method)
	params.Put(VERSION, req.Version)
	params.Put(APP_ID, a.AppId)
	params.Put(SIGN_TYPE, a.SignType)
	params.Put(NOTIFY_URL, req.NotifyUrl)
	params.Put(RETURN_URL, req.ReturnUrl)
	params.Put(CHARSET, a.Charset)
	params.Put(FORMAT, a.Format)
	params.Put(SIGN_TYPE, a.SignType)
	params.Put(TIMESTAMP, time.Now().Format(DATE_TIME_FORMAT))
	bs, err := json.Marshal(req.BizModel)
	if err != nil {
		return nil, err
	}
	params.Put(BIZ_CONTENT_KEY, string(bs))
	return params, nil
}

func (a *AlipayClient) signParams(params AlipayParams) error {
	keys := make([]string, 0, len(params))
	for key, _ := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var buffer bytes.Buffer
	for i := 0; i < len(keys); i++ {
		if i > 0 {
			buffer.WriteString("&")
		}
		buffer.WriteString(keys[i] + "=" + params[keys[i]])
	}
	sign, err := a.sign(buffer.Bytes())
	if err != nil {
		return err
	} else {
		params.Put(SIGN, sign)
		return nil
	}
}

func (a *AlipayClient) sign(data []byte) (string, error) {
	block, _ := pem.Decode([]byte(a.PrivateKey))
	if block == nil {
		return "", errors.New("[gopay.alipay] private key pem decode error")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", errors.New("[gopay.alipay] private key parse error, " + err.Error())
	}

	var (
		h          hash.Hash
		cryptoHash crypto.Hash
	)

	if a.SignType == SIGN_TYPE_RSA {
		h = sha1.New()
		cryptoHash = crypto.SHA1
	} else {
		h = sha256.New()
		cryptoHash = crypto.SHA256
	}
	h.Write(data)
	digest := h.Sum(nil)
	s, err := rsa.SignPKCS1v15(nil, privateKey, cryptoHash, digest)
	if err != nil {
		return "", errors.New("[gopay.alipay] sign error, " + err.Error())
	}
	return base64.StdEncoding.EncodeToString(s), nil
}

func (a *AlipayClient) VerifySign(content []byte, sign string) error {
	sig, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return errors.New("[gopay.alipay] sign decode error, " + err.Error())
	}
	block, _ := pem.Decode([]byte(a.AlipayPublicKey))
	if block == nil {
		return errors.New("[gopay.alipay] alipay public key pem decode error")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return errors.New("[gopay.alipay] alipay public key parse error, " + err.Error())
	}

	var (
		h          hash.Hash
		cryptoHash crypto.Hash
	)

	if a.SignType == SIGN_TYPE_RSA {
		h = sha1.New()
		cryptoHash = crypto.SHA1
	} else {
		h = sha256.New()
		cryptoHash = crypto.SHA256
	}

	h.Write(content)
	digest := h.Sum(nil)
	return rsa.VerifyPKCS1v15(pub.(*rsa.PublicKey), cryptoHash, digest, sig)
}

func (a *AlipayClient) parseResponse(r io.Reader, req *AlipayRequest) error {
	decoder := json.NewDecoder(r)
	var (
		sign       string
		contentKey string
		content    json.RawMessage
	)
	contentKey = strings.Replace(req.method, ".", "_", -1) + RESPONSE_SUFFIX
	for {
		t, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		switch t.(type) {
		case string:
			if t.(string) == SIGN {
				decoder.Decode(&sign)
			} else if t.(string) == contentKey {
				decoder.Decode(&content)
			} else {
				var raw json.RawMessage
				decoder.Decode(&raw)
			}
		default:
			continue
		}
	}

	if err := a.VerifySign(content, sign); err != nil {
		if err := a.VerifySign(bytes.Replace(content, []byte("\\/"), []byte("/"), -1), sign); err != nil {
			return errors.New("[alipay] sign check fail: check Sign and Data Fail！JSON also！")
		}
	}

	return json.Unmarshal(content, req.response)
}
