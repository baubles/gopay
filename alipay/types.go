package alipay

type AlipayParams map[string]string

func MakeAlipayParams() AlipayParams {
	return make(AlipayParams)
}

func (a AlipayParams) Put(key, value string) {
	a[key] = value
}

func (a AlipayParams) Get(key string) string {
	v, _ := a[key]
	return v
}

func (a AlipayParams) Has(key string) bool {
	_, ok := a[key]
	return ok
}

const (
	SIGN_TYPE                      = "sign_type"
	SIGN_TYPE_RSA                  = "RSA"
	SIGN_TYPE_RSA2                 = "RSA2"
	SIGN_ALGORITHMS                = "SHA1WithRSA"
	SIGN_SHA256RSA_ALGORITHMS      = "SHA256WithRSA"
	ENCRYPT_TYPE_AES               = "AES"
	APP_ID                         = "app_id"
	FORMAT                         = "format"
	METHOD                         = "method"
	TIMESTAMP                      = "timestamp"
	VERSION                        = "version"
	SIGN                           = "sign"
	ALIPAY_SDK                     = "alipay_sdk"
	ACCESS_TOKEN                   = "auth_token"
	APP_AUTH_TOKEN                 = "app_auth_token"
	TERMINAL_TYPE                  = "terminal_type"
	TERMINAL_INFO                  = "terminal_info"
	CHARSET                        = "charset"
	NOTIFY_URL                     = "notify_url"
	RETURN_URL                     = "return_url"
	ENCRYPT_TYPE                   = "encrypt_type"
	BIZ_CONTENT_KEY                = "biz_content"
	DATE_TIME_FORMAT               = "2006-01-02 15:04:05"
	DATE_TIMEZONE                  = "GMT+8"
	CHARSET_UTF8                   = "UTF-8"
	CHARSET_GBK                    = "GBK"
	FORMAT_JSON                    = "json"
	FORMAT_XML                     = "xml"
	PROD_CODE                      = "prod_code"
	ERROR_RESPONSE                 = "error_response"
	RESPONSE_SUFFIX                = "_response"
	RESPONSE_XML_ENCRYPT_NODE_NAME = "response_encrypted"
)

// 商品详情
type GoodsDetail struct {
	GoodsId       string `json:"goods_id"`       // 商品的编号
	GoodsName     string `json:"goods_name"`     // 商品名称
	Quantity      string `json:"quantity"`       // 商品数量
	Price         string `json:"price"`          // 商品单价，单位为元
	GoodsCategory string `json:"goods_category"` // 商品类目
	Body          string `json:"body"`           // 商品描述信息
	ShowUrl       string `json:"show_url"`       // 商品的展示地址
}

type TradeFundBill struct {
	FundChannel string `json:"fund_channel"`
	Amount      string `json:"amount"`
	RealAmount  string `json:"real_amount"`
	FundType    string `json:"fund_type"`
}
