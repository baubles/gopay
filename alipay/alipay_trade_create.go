package alipay

// 下单参数 参考: https://docs.open.alipay.com/api_1/alipay.trade.create/
type AlipayTradeCreateBizModel struct {
	OutTradeNo         string        `json:"out_trade_no"`
	SellerId           string        `json:"seller_id"`
	TotalAmount        string        `json:"total_amount"`
	DiscountableAmount string        `json:"discountable_amount"`
	Subject            string        `json:"subject"`
	Body               string        `json:"body"`
	BuyerId            string        `json:"buyer_id"`
	GoodsDetail        []GoodsDetail `json:"goods_detail"`
	OperatorId         string        `json:"operator_id"`
	StoreId            string        `json:"store_id"`
	TerminalId         string        `json:"terminal_id"`
	ExtendParams       string        `json:"extend_params"`
	TimeoutExpress     string        `json:"timeout_express"`
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
