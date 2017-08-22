package wxpay

import (
	"bytes"
	"encoding/xml"
	"io"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
)

const (
	APP_ID                = "appid"
	MCH_ID                = "mch_id"
	SUB_APP_ID            = "sub_appid"
	SUB_MCH_ID            = "sub_mch_id"
	SIGN                  = "sign"
	SIGN_TYPE             = "sign_type"
	SIGN_TYPE_MD5         = "MD5"
	SIGN_TYPE_HMAC_SHA256 = "HMAC-SHA256"
	TRADE_TYPE_APP        = "APP"
	NONCE_STR             = "nonce_str"
)

type WXPayRequest struct {
	apiUrl         string
	requireCert    bool
	bizModel       interface{}
	responseParser WXPayParamsParser
}

type WXPayResponse struct {
	ReturnCode string `ApiField:"return_code"` // 返回状态码
	ReturnMsg  string `ApiField:"return_msg"`  // 返回信息

	// 以下字段在return_code为SUCCESS的时候有返回
	Appid      string `ApiField:"appid"`        // 应用APPID
	MchId      string `ApiField:"mch_id"`       // 商户号
	DeviceInfo string `ApiField:"device_info"`  // 设备号
	NonceStr   string `ApiField:"nonce_str"`    // 随机字符串
	Sign       string `ApiField:"sign"`         // 签名
	ResultCode string `ApiField:"result_code"`  // 业务结果
	ErrCode    string `ApiField:"err_code"`     // 错误代码
	ErrCodeDes string `ApiField:"err_code_des"` // 错误代码描述
}

type WXPayParams map[string]string

func (self WXPayParams) Put(key, value string) {
	if len(strings.TrimSpace(value)) > 0 {
		self[key] = value
	}

}

func (self WXPayParams) Get(key string) string {
	v, _ := self[key]
	return v
}

func (self WXPayParams) Has(key string) bool {
	_, ok := self[key]
	return ok
}

func (self WXPayParams) Xml() []byte {
	buf := new(bytes.Buffer)
	encoder := xml.NewEncoder(buf)
	xmlStart := xml.StartElement{Name: xml.Name{Local: "xml"}}
	encoder.Indent("", "  ")
	encoder.EncodeToken(xmlStart)
	for k, v := range self {
		token := xml.StartElement{Name: xml.Name{Local: k}}
		// encoder.EncodeElement(struct {
		// 	string `xml:",cdata"`
		// }{v}, token)
		encoder.EncodeElement(v, token)
	}
	encoder.EncodeToken(xmlStart.End())
	encoder.Flush()
	return buf.Bytes()
}

func (self WXPayParams) Parse(r io.Reader) error {
	var start *xml.StartElement
	d := xml.NewDecoder(r)
	for {
		t, err := d.Token()
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
		switch t.(type) {
		case xml.StartElement:
			s := t.(xml.StartElement)
			start = &s
			break
		case xml.CharData:
			c := t.(xml.CharData)
			if c = bytes.TrimSpace(c); len(c) > 0 {
				self[start.Name.Local] = string(c)
			}
			break
		}
	}
	return nil
}

func (self WXPayParams) toStruct(val reflect.Value) {
	if val.Type().Kind() == reflect.Ptr && val.IsNil() {
		val.Set(reflect.New(val.Elem().Type()))
	}
	ind := reflect.Indirect(val)
	typ := ind.Type()
	for i := 0; i < ind.NumField(); i++ {
		fVal := ind.Field(i)
		fTyp := typ.Field(i)
		if fTyp.Anonymous && fTyp.Type.Kind() == reflect.Struct {
			self.toStruct(fVal)
		}
		if name, ok := fTyp.Tag.Lookup("ApiField"); ok && fTyp.Type.Kind() == reflect.String {
			if v, ok := self[name]; ok {
				fVal.SetString(v)
			}
		}
	}
}

func nonceStr() string {
	return strconv.Itoa(rand.Int())
}

func MakeWXPayParams() WXPayParams {
	return make(WXPayParams)
}

type WXPayParamsParser interface {
	Parse(params WXPayParams)
}

// 微信代金券
type WXPayCoupon struct {
	CouponId   string // 代金券id
	CouponType string // 代金券类型
	CouponFee  string // 代金券金额
}

type WXPayRefund struct {
	OutRefundNo       string // 商户退款单号
	RefundId          string // 微信退款单号
	RefundChannel     string // 退款渠道
	RefundFee         string // 退款金额
	CouponRefundFee   string // 代金券退款金额
	CouponRefundCount string // 代金券使用数量
	RefundStatus      string // 退款状态
	RefundRecvAccout  string // 退款入账账户
	RefundSuccessTime string // 退款成功时间
	Coupons           []WXPayCoupon
}
