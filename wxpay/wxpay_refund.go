package wxpay

import (
	"reflect"
	"strconv"
)

// 微信支付退款请求
type WXPayRefundBizModel struct {
	SignType      string `ApiField:"sign_type"`       // 签名类型
	TransactionId string `ApiField:"transaction_id"`  // 微信订单号
	OutTradeNo    string `ApiField:"out_trade_no"`    // 商户订单号
	OutRefundNo   string `ApiField:"out_refund_no"`   // 商户退款单号
	TotalFee      string `ApiField:"total_fee"`       // 订单金额
	RefundFee     string `ApiField:"refund_fee"`      // 退款金额
	RefundFeeType string `ApiField:"refund_fee_type"` // 货币种类
	RefundDesc    string `ApiField:"refund_desc"`     // 退款原因
	RefundAccount string `ApiField:"refund_account"`  // 退款资金来源
}

// 微信支付退款返回
type WXPayRefundResponse struct {
	WXPayResponse

	TransactionId       string `ApiField:"transaction_id"`        // 微信订单号
	OutTradeNo          string `ApiField:"out_trade_no"`          // 商户订单号
	OutRefundNo         string `ApiField:"out_refund_no"`         // 商户退款单号
	RefundId            string `ApiField:"refund_id"`             // 微信退款单号
	RefundFee           string `ApiField:"refund_fee"`            // 退款金额
	SettlementRefundFee string `ApiField:"settlement_refund_fee"` // 应结退款金额
	TotalFee            string `ApiField:"total_fee"`             // 标价金额
	SettlementTotalFee  string `ApiField:"settlement_total_fee"`  // 应结订单金额
	FeeType             string `ApiField:"fee_type"`              // 标价币种
	CashFee             string `ApiField:"cash_fee"`              // 现金支付金额
	CashFeeType         string `ApiField:"cash_fee_type"`         // 现金支付币种
	CashRefundFee       string `ApiField:"cash_refund_fee"`       // 现金退款金额
	CouponRefundFee     string `ApiField:"coupon_refund_fee"`     // 代金券退款总金额
	CouponRefundCount   string `ApiField:"coupon_refund_count"`   // 退款代金券使用数量
	CouponRefunds       []WXPayCoupon
}

func (self *WXPayRefundResponse) Parse(params WXPayParams) {
	params.toStruct(reflect.ValueOf(self))

	// 处理代金券
	count, _ := strconv.Atoi(self.CouponRefundCount)
	coupons := make([]WXPayCoupon, count)
	for i := 0; i < count; i++ {
		coupons[i] = WXPayCoupon{
			CouponId:   params["coupon_id_"+strconv.Itoa(i)],
			CouponType: params["coupon_type_"+strconv.Itoa(i)],
			CouponFee:  params["coupon_fee_"+strconv.Itoa(i)],
		}
	}
	self.CouponRefunds = coupons
}

func NewWXPayRefundRequest(bizModel WXPayRefundBizModel) *WXPayRequest {
	return &WXPayRequest{
		apiUrl:         "https://api.mch.weixin.qq.com/secapi/pay/refund",
		requireCert:    true,
		bizModel:       bizModel,
		responseParser: new(WXPayRefundResponse),
	}
}
