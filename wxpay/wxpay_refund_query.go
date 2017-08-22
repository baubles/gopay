package wxpay

import (
	"reflect"
	"strconv"
)

type WXPayRefundQueryBizModel struct {
	SubAppid      string `ApiField:"sub_appid"`      // 子商户应用ID
	SubMchId      string `ApiField:"sub_mch_id"`     // 子商户号
	TransactionId string `ApiField:"transaction_id"` // 微信订单号
	OutTradeNo    string `ApiField:"out_trade_no"`   // 商户订单号
	OutRefundNo   string `ApiField:"out_refund_no"`  // 商户退款单号
	RefundId      string `ApiField:"refund_id"`      // 微信退款单号
}

//  退款结果返回
type WXPayRefundQueryResponse struct {
	WXPayResponse

	SubAppid      string `ApiField:"sub_appid"`      // 子商户应用ID
	SubMchId      string `ApiField:"sub_mch_id"`     // 子商户号
	TransactionId string `ApiField:"transaction_id"` // 微信订单号
	OutTradeNo    string `ApiField:"out_trade_no"`   // 商户订单号
	TotalFee      string `ApiField:"total_fee"`      // 订单总金额
	FeeType       string `ApiField:"fee_type"`       // 订单金额货币种类
	CashFee       string `ApiField:"cash_fee"`       // 现金支付金额
	RefundCount   string `ApiField:"refund_count"`   // 退款笔数
	Refunds       []WXPayRefund
}

func (self *WXPayRefundQueryResponse) Parse(params WXPayParams) {
	params.toStruct(reflect.ValueOf(self))

	// 处理退款
	count, _ := strconv.Atoi(self.RefundCount)
	refunds := make([]WXPayRefund, count)
	for i := 0; i < count; i++ {
		refunds[i] = WXPayRefund{
			OutRefundNo:       params["out_refund_no_"+strconv.Itoa(i)],
			RefundId:          params["refund_id_"+strconv.Itoa(i)],
			RefundChannel:     params["refund_channel_"+strconv.Itoa(i)],
			RefundFee:         params["refund_fee_"+strconv.Itoa(i)],
			CouponRefundFee:   params["coupon_refund_fee_"+strconv.Itoa(i)],
			CouponRefundCount: params["coupon_refund_count_"+strconv.Itoa(i)],
			RefundStatus:      params["refund_status_"+strconv.Itoa(i)],
			RefundRecvAccout:  params["refund_recv_accout_"+strconv.Itoa(i)],
			RefundSuccessTime: params["refund_success_time_"+strconv.Itoa(i)],
		}

		if coupNum, err := strconv.Atoi(refunds[i].CouponRefundCount); err != nil {
			coupons := make([]WXPayCoupon, coupNum)
			for j := 0; j < coupNum; j++ {
				coupons[j] = WXPayCoupon{
					CouponId:  params["coupon_refund_id_"+strconv.Itoa(i)+"_"+strconv.Itoa(j)],
					CouponFee: params["coupon_refund_fee_"+strconv.Itoa(i)+"_"+strconv.Itoa(j)],
				}
			}
		}
	}

	self.Refunds = refunds
}

func NewWXRefundQueryRequest(bizModel WXPayOrderQueryBizModel) *WXPayRequest {
	return &WXPayRequest{
		requireCert:    false,
		apiUrl:         "https://api.mch.weixin.qq.com/pay/refundquery",
		bizModel:       bizModel,
		responseParser: new(WXPayRefundQueryResponse),
	}
}
