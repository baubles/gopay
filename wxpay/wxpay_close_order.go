package wxpay

import (
	"reflect"
)

// 微信关闭订单请求
type WXPayCloseOrderBizModel struct {
	OutTradeNo string `ApiField:"out_trade_no"` // 商户订单号
}

// 微信关闭订单返回
type WXPayCloseOrderResponse struct {
	WXPayResponse
}

func (self *WXPayCloseOrderResponse) Parse(params WXPayParams) {
	params.toStruct(reflect.ValueOf(self))
}

func NewWXPayCloseOrderRequest(bizModel WXPayCloseOrderBizModel) *WXPayRequest {
	return &WXPayRequest{
		apiUrl:         "https://api.mch.weixin.qq.com/pay/closeorder",
		requireCert:    false,
		bizModel:       bizModel,
		responseParser: new(WXPayCloseOrderResponse),
	}
}
