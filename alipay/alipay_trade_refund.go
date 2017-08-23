package alipay

// 退款参数
type AlipayTradeRefundBizModel struct {
	OutTradeNo   string `json:"out_trade_no,omitempty"`   // 订单支付时传入的商户订单号,不能和 trade_no同时为空。
	TradeNo      string `json:"trade_no,omitempty"`       // 支付宝交易号，和商户订单号不能同时为空
	RefundAmount string `json:"refund_amount,omitempty"`  // 需要退款的金额，该金额不能大于订单金额,单位为元，支持两位小数
	RefundReason string `json:"refund_reason,omitempty"`  // 退款的原因说明
	OutRequestNo string `json:"out_request_no,omitempty"` // 标识一次退款请求，同一笔交易多次退款需要保证唯一，如需部分退款，则此参数必传。
	OperatorId   string `json:"operator_id,omitempty"`    // 商户的操作员编号
	StoreId      string `json:"store_id,omitempty"`       // 商户的门店编号
	TerminalId   string `json:"terminal_id,omitempty"`    // 商户的终端编号
}

type AlipayTradeRefundResponse struct {
	AlipayResponse
	TradeNo              string          `json:"trade_no"`
	OutTradeNo           string          `json:"out_trade_no"`
	BuyerLogonId         string          `json:"buyer_logon_id"`
	FundChange           string          `json:"fund_change"`
	RefundFee            string          `json:"refund_fee"`
	GmtRefundPay         string          `json:"gmt_refund_pay"`
	RefundDetailItemList []TradeFundBill `json:"refund_detail_item_list"`
	StoreName            string          `json:"store_name"`
	BuyerUserId          string          `json:"buyer_user_id"`
}

func NewAlipayTradeRefundRequest(bizModel AlipayTradeRefundBizModel) *AlipayRequest {
	return &AlipayRequest{
		method:   "alipay.trade.refund",
		BizModel: bizModel,
		response: new(AlipayTradeRefundResponse),
		Version:  "1.0",
	}
}
