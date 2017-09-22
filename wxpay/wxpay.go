package wxpay

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"sort"
	"time"
)

var Debug bool = false

type WXPayClient struct {
	Appid      string
	MchId      string
	PrivateKey string
	SubAppid   string
	SubMchId   string
	SignType   string
	isCertOk   bool
	httpClient *http.Client
}

func NewWXPayClient(appid, mchId, privateKey string) *WXPayClient {
	return &WXPayClient{
		Appid:      appid,
		MchId:      mchId,
		PrivateKey: privateKey,
		SignType:   SIGN_TYPE_MD5,
		httpClient: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

func (self *WXPayClient) SetTslCert(certPEMBlock, keyPEMBlock, rootCAPEMBlock []byte) error {
	certificate, err := tls.X509KeyPair(certPEMBlock, keyPEMBlock)
	if err != nil {
		return err
	}
	rootCAs := x509.NewCertPool()
	if ok := rootCAs.AppendCertsFromPEM(rootCAPEMBlock); !ok {
		return errors.New("[gopay.wxpay] root ca pem is error")
	}
	self.httpClient.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:      rootCAs,
			Certificates: []tls.Certificate{certificate},
		},
	}
	self.isCertOk = true
	return nil
}

func (self *WXPayClient) Do(req *WXPayRequest) (interface{}, error) {
	var err error
	params := self.params(req)
	self.SignParams(params)
	if req.requireCert && !self.isCertOk {
		return nil, errors.New("[gopay.wxpay] request require cert")
	}
	resp, err := self.httpClient.Post(req.apiUrl, "text/xml", bytes.NewReader(params.Xml()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respParams := MakeWXPayParams()
	body, err := ioutil.ReadAll(resp.Body)
	err = respParams.Parse(bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	if Debug {
		log.Println(req.apiUrl, string(params.Xml()), respParams)
	}
	if !self.checkSign(params) {
		return nil, errors.New("[gopay.wxpay] response sign error")
	}
	req.responseParser.Parse(respParams)
	return req.responseParser, nil
}

func (self *WXPayClient) params(req *WXPayRequest) WXPayParams {
	params := MakeWXPayParams()
	params.Put(APP_ID, self.Appid)
	params.Put(MCH_ID, self.MchId)
	params.Put(SUB_APP_ID, self.SubAppid)
	params.Put(SUB_MCH_ID, self.SubMchId)
	params.Put(SIGN_TYPE, self.SignType)
	params.Put(NONCE_STR, nonceStr())
	bVal := reflect.ValueOf(req.bizModel)
	bInd := reflect.Indirect(bVal)
	bTyp := bInd.Type()
	for i := 0; i < bInd.NumField(); i++ {
		field := bInd.Field(i)
		fType := bTyp.Field(i)
		if apiField := fType.Tag.Get("ApiField"); apiField != "" && field.IsValid() && field.Kind() == reflect.String {
			params.Put(apiField, field.String())
		}
	}
	return params
}

func (self *WXPayClient) checkSign(params WXPayParams) bool {
	sign := params.Get(SIGN)
	self.SignParams(params)
	if sign == params.Get(SIGN) {
		return true
	}
	return false
}

func (self *WXPayClient) SignParams(params WXPayParams) {
	keys := make([]string, 0, len(params))
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var buffer bytes.Buffer
	for i := 0; i < len(keys); i++ {
		if keys[i] == SIGN {
			continue
		}
		if i > 0 {
			buffer.WriteString("&")
		}
		buffer.WriteString(keys[i] + "=" + params[keys[i]])
	}
	buffer.WriteString("&key=" + self.PrivateKey)
	if params.Get(SIGN_TYPE) == SIGN_TYPE_HMAC_SHA256 {
		params.Put(SIGN, self.hmacSha256(buffer.Bytes()))
	} else {
		params.Put(SIGN, self.md5(buffer.Bytes()))
	}

}

func (self *WXPayClient) hmacSha256(data []byte) string {
	// TODO 退款接口会签名错误
	h := hmac.New(sha256.New, []byte(self.PrivateKey))
	h.Write(data)
	return fmt.Sprintf("%X", h.Sum(nil))
}

func (self *WXPayClient) md5(data []byte) string {
	h := md5.New()
	h.Write(data)
	return fmt.Sprintf("%X", h.Sum(nil))
}
