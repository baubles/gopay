package alipay

type AlipayTradeQueryBizModel struct {
	TradeNo    string `json:"trade_no,omitempty"`     // 商户订单号
	OutTradeNo string `json:"out_trade_no,omitempty"` // 支付宝交易号
}

type AlipayTradeQueryResponse struct {
	AlipayResponse
	TradeNo        string          `json:"trade_no"`
	OutTradeNo     string          `json:"out_trade_no"`
	BuyerLogonId   string          `json:"buyer_logon_id"`
	TradeStatus    string          `json:"trade_status"`
	TotalAmount    string          `json:"total_amount"`
	ReceiptAmount  string          `json:"receipt_amount"`
	BuyerPayAmount string          `json:"buyer_pay_amount"`
	PointAmount    string          `json:"point_amount"`
	InvoiceAmount  string          `json:"invoice_amount"`
	SendPayDate    string          `json:"send_pay_date"`
	StoreId        string          `json:"store_id"`
	TerminalId     string          `json:"terminal_id"`
	FundBillList   []TradeFundBill `json:"fund_bill_list"`
	StoreName      string          `json:"store_name"`
	BuyerUserId    string          `json:"buyer_user_id"`
}

func NewAlipayTradeQueryRequest(bizModel AlipayTradeQueryBizModel) *AlipayRequest {
	return &AlipayRequest{
		method:   "alipay.trade.query",
		BizModel: bizModel,
		response: new(AlipayTradeQueryResponse),
		Version:  "1.0",
	}
}
