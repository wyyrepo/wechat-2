package wechat

import "encoding/xml"

// 交易保障(MICROPAY)
func (c *Client) ReportMicropay(body ReportMicropayBody) (wxRsp ReportMicropayResponse, err error) {
	// 处理参数
	if body.InterfaceUrl, err = EscapedPath("https://api.mch.weixin.qq.com/pay/batchreport/micropay/total"); err != nil {
		return
	}
	body.TradesStr = JsonString(body.Trades)
	// 业务逻辑
	bytes, err := c.doWeChat("payitil/report", body)
	if err != nil {
		return
	}
	// 解析返回值
	err = xml.Unmarshal(bytes, &wxRsp)
	return
}

// 交易保障(MICROPAY)的参数
type ReportMicropayBody struct {
	SignType     string `json:"sign_type,omitempty"`   // 签名类型，目前支持HMAC-SHA256和MD5，默认为MD5
	DeviceInfo   string `json:"device_info,omitempty"` // (非必填) 微信支付分配的终端设备号，商户自定义
	InterfaceUrl string `json:"interface_url"`         // (不需要手动填写) 上报对应的接口的完整URL，类似：https://api.mch.weixin.qq.com/pay/unifiedorder 对于刷卡支付，为更好的和商户共同分析一次业务行为的整体耗时情况，对于两种接入模式，请都在门店侧对一次刷卡行为进行一次单独的整体上报，上报URL指定为：https://api.mch.weixin.qq.com/pay/micropay/total 关于两种接入模式具体可参考本文档章节：刷卡支付商户接入模式 其它接口调用仍然按照调用一次，上报一次来进行。
	UserIp       string `json:"user_ip"`               // 发起接口调用时的机器IP
	TradesStr    string `json:"trades"`                // POS机采集的交易信息列表，使用JSON格式的数组
	// 生成TradesStr
	Trades []ReportMicropayBodyTrade `json:"-"`
}

type ReportMicropayBodyTrade struct {
	OutTradeNo string `json:"out_trade_no"`      // 商户订单号
	BeginTime  string `json:"begin_time"`        // 交易开始时间(扫码时间)
	EndTime    string `json:"end_time"`          // 交易完成时间
	State      string `json:"state"`             // 交易结果，OK-成功 FAIL-失败 CANCLE-取消
	ErrMsg     string `json:"err_msg,omitempty"` // 自定义的错误描述信息
}

// 交易保障(MICROPAY)的返回值
type ReportMicropayResponse struct {
	ResponseModel
	// 当return_code为SUCCESS时
	ResultCode string `xml:"result_code"` // SUCCESS/FAIL
}
