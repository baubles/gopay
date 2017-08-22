package wxpay

import (
	"reflect"
	"strconv"
)

// 微信订单查询请求参数
type WXPayOrderQueryBizModel struct {
	TransactionId string `ApiField:"transaction_id"` // 微信订单号
	OutTradeNo    string `ApiField:"out_trade_no"`   // 商户订单号
}

// 微信订单查询结果返回
type WXPayOrderQueryResponse struct {
	WXPayResponse

	// 以下字段在return_code 和result_code都为SUCCESS的时候有返回
	DeviceInfo         string `ApiField:"device_info"`          //设备号
	Openid             string `ApiField:"openid"`               //用户标识
	IsSubscribe        string `ApiField:"is_subscribe"`         //是否关注公众账号
	TradeType          string `ApiField:"trade_type"`           //交易类型
	TradeState         string `ApiField:"trade_state"`          //交易状态
	BankType           string `ApiField:"bank_type"`            //付款银行
	TotalFee           string `ApiField:"total_fee"`            //总金额
	FeeType            string `ApiField:"fee_type"`             //货币种类
	CashFee            string `ApiField:"cash_fee"`             //现金支付金额
	CashFeeType        string `ApiField:"cash_fee_type"`        //现金支付货币类型
	SettlementTotalFee string `ApiField:"settlement_total_fee"` //应结订单金额
	CouponFee          string `ApiField:"coupon_fee"`           //代金券金额
	CouponCount        string `ApiField:"coupon_count"`         //代金券使用数量
	TransactionId      string `ApiField:"transaction_id"`       //微信支付订单号
	OutTradeNo         string `ApiField:"out_trade_no"`         //商户订单号
	Attach             string `ApiField:"attach"`               //附加数据
	TimeEnd            string `ApiField:"time_end"`             //支付完成时间
	TradeStateDesc     string `ApiField:"trade_state_desc"`     //交易状态描述

	// coupon_id_$n string `xml:"coupon_id_$n"` //代金券ID
	// coupon_type_$n string `xml:"coupon_type_$n"` //代金券类型
	// coupon_fee_$n string `xml:"coupon_fee_$n"` //单个代金券支付金额
	Coupons []WXPayCoupon // 代金券列表
}

func (self *WXPayOrderQueryResponse) Parse(params WXPayParams) {
	params.toStruct(reflect.ValueOf(self))

	// 处理代金券
	count, _ := strconv.Atoi(self.CouponCount)
	coupons := make([]WXPayCoupon, count)
	for i := 0; i < count; i++ {
		coupons[i] = WXPayCoupon{
			CouponId:   params["coupon_id_"+strconv.Itoa(i)],
			CouponType: params["coupon_type_"+strconv.Itoa(i)],
			CouponFee:  params["coupon_fee_"+strconv.Itoa(i)],
		}
	}
	self.Coupons = coupons
}

func NewWXPayOrderQueryRequest(bizModel WXPayOrderQueryBizModel) *WXPayRequest {
	return &WXPayRequest{
		apiUrl:         "https://api.mch.weixin.qq.com/pay/orderquery",
		requireCert:    false,
		bizModel:       bizModel,
		responseParser: new(WXPayOrderQueryResponse),
	}
}
