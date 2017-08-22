package wxpay

import (
	"reflect"
)

// 微信统一下单参数
type WXPayUnifiedOrderBizModel struct {
	DeviceInfo     string `ApiField:"device_info"`      // 设备号
	Body           string `ApiField:"body"`             // 商品描述
	Detail         string `ApiField:"detail"`           // 商品详情
	Attach         string `ApiField:"attach"`           // 附加数据
	OutTradeNo     string `ApiField:"out_trade_no"`     // 商户订单号
	FeeType        string `ApiField:"fee_type"`         // 货币类型
	TotalFee       string `ApiField:"total_fee"`        // 总金额
	SpbillCreateIp string `ApiField:"spbill_create_ip"` // 终端IP
	TimeStart      string `ApiField:"time_start"`       // 交易起始时间
	TimeExpire     string `ApiField:"time_expire"`      // 交易结束时间
	GoodsTag       string `ApiField:"goods_tag"`        // 订单优惠标记
	NotifyUrl      string `ApiField:"notify_url"`       // 通知地址
	TradeType      string `ApiField:"trade_type"`       // 交易类型
	LimitPay       string `ApiField:"limit_pay"`        // 指定支付方式
	SceneInfo      string `ApiField:"scene_info"`       // 场景信息
}

// 微信统一下单返回结果
type WXPayUnifiedOrderResponse struct {
	WXPayResponse

	// 以下字段在return_code 和result_code都为SUCCESS的时候有返回
	TradeType string `ApiField:"trade_type"` // 交易类型
	PrepayId  string `ApiField:"prepay_id"`  // 预支付交易会话标识
}

func (self *WXPayUnifiedOrderResponse) Parse(params WXPayParams) {
	params.toStruct(reflect.ValueOf(self))
}

func NewWXPayUnifiedOrderRequest(bizModel WXPayUnifiedOrderBizModel) *WXPayRequest {
	return &WXPayRequest{
		apiUrl:         "https://api.mch.weixin.qq.com/pay/unifiedorder",
		requireCert:    false,
		bizModel:       bizModel,
		responseParser: new(WXPayUnifiedOrderResponse),
	}
}
