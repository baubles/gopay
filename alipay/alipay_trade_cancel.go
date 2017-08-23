package alipay

// 取消订单参数
type AlipayTradeCancelBizModel struct {
	TradeNo    string `json:"trade_no,omitempty"`     // 商户订单号
	OutTradeNo string `json:"out_trade_no,omitempty"` // 支付宝交易号
}

type AlipayTradeCancelResponse struct {
	AlipayResponse
	TradeNo    string `json:"trade_no"`
	OutTradeNo string `json:"out_trade_no"`
	RetryFlag  string `json:"retry_flag"`
	Action     string `json:"action"`
}

func NewAlipayTradeCancelRequest(bizModel AlipayTradeCancelBizModel) *AlipayRequest {
	return &AlipayRequest{
		method:   "alipay.trade.cancel",
		BizModel: &bizModel,
		response: new(AlipayTradeCancelResponse),
		Version:  "1.0",
	}
}
