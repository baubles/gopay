package alipay

// 下单参数 参考: https://docs.open.alipay.com/api_1/alipay.trade.create/
type AlipayTradeCreateBizModel struct {
	OutTradeNo         string        `json:"out_trade_no,omitempty"`
	SellerId           string        `json:"seller_id,omitempty"`
	TotalAmount        string        `json:"total_amount,omitempty"`
	DiscountableAmount string        `json:"discountable_amount,omitempty"`
	Subject            string        `json:"subject,omitempty"`
	Body               string        `json:"body,omitempty"`
	BuyerId            string        `json:"buyer_id,omitempty"`
	GoodsDetail        []GoodsDetail `json:"goods_detail,omitempty"`
	OperatorId         string        `json:"operator_id,omitempty"`
	StoreId            string        `json:"store_id,omitempty"`
	TerminalId         string        `json:"terminal_id,omitempty"`
	ExtendParams       string        `json:"extend_params,omitempty"`
	TimeoutExpress     string        `json:"timeout_express,omitempty"`
}

type AlipayTradeCreateResponse struct {
	AlipayResponse
	OutTradeNo string `json:"out_trade_no"` // 商户订单号
	TradeNo    string `json:"trade_no"`     //支付宝交易号
}

func NewAlipayTradeCreateRequest(bizModel AlipayTradeCreateBizModel) *AlipayRequest {
	return &AlipayRequest{
		method:   "alipay.trade.create",
		BizModel: &bizModel,
		response: new(AlipayResponse),
		Version:  "1.0",
	}
}
