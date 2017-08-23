package alipay

// 关闭订单参数
type AlipayTradeCloseBizModel struct {
	TradeNo    string `json:"trade_no,omitempty"`     // 商户订单号
	OutTradeNo string `json:"out_trade_no,omitempty"` // 支付宝交易号
	OperatorId string `json:"operator_id,omitempty"`  //
}

type AlipayTradeCloseResponse struct {
	AlipayResponse
	OutTradeNo string `json:"out_trade_no"` // 商户订单号
	TradeNo    string `json:"trade_no"`     // 支付宝交易号
}

func NewAlipayTradeCloseRequest(bizModel AlipayTradeCloseBizModel) *AlipayRequest {
	return &AlipayRequest{
		method:   "alipay.trade.close",
		BizModel: bizModel,
		response: new(AlipayTradeCloseResponse),
		Version:  "1.0",
	}
}
