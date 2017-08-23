package alipay

type AlipayTradeAppPayBizModel struct {
	Body           string `json:"body,omitempty"`
	Subject        string `json:"subject,omitempty"`
	OutTradeNo     string `json:"out_trade_no,omitempty"`
	TimeoutExpress string `json:"timeout_express,omitempty"`
	TotalAmount    string `json:"total_amount,omitempty"`
	ProductCode    string `json:"product_code,omitempty"`
	GoodsType      string `json:"goods_type,omitempty"`
	PassbackParams string `json:"passback_params,omitempty"`
	PromoParams    string `json:"promo_params,omitempty"`
	ExtendParams   struct {
		hbFqNum              string `json:"hb_fq_num,omitempty"`
		hbFqSellerPercent    string `json:"hb_fq_seller_percent,omitempty"`
		sysServiceProviderId string `json:"sys_service_provider_id,omitempty"`
	} `json:"extend_params,omitempty"`
	EnablePayChannels  string `json:"enable_pay_channels,omitempty"`
	DisablePayChannels string `json:"disable_pay_channels,omitempty"`
	StoreId            string `json:"store_id,omitempty"`
}

func NewAlipayTradeAppPayRequest(bizModel AlipayTradeAppPayBizModel) *AlipayRequest {
	return &AlipayRequest{
		method:   "alipay.trade.app.pay",
		BizModel: &bizModel,
		Version:  "1.0",
		response: new(AlipayResponse),
	}
}
