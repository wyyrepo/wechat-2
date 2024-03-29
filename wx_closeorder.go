package wechat

import "encoding/xml"

// 关闭订单
func (c *Client) CloseOrder(body CloseOrderBody) (wxRsp CloseOrderResponse, err error) {
	// 业务逻辑
	bytes, err := c.doWeChat("pay/closeorder", body)
	if err != nil {
		return
	}
	// 结果校验
	if err = c.doVerifySign(bytes, true); err != nil {
		return
	}
	// 解析返回值
	err = xml.Unmarshal(bytes, &wxRsp)
	return
}

// 关闭订单的参数
type CloseOrderBody struct {
	SignType   string `json:"sign_type,omitempty"` // 签名类型，目前支持HMAC-SHA256和MD5，默认为MD5
	OutTradeNo string `json:"out_trade_no"`        // 商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*且在同一个商户号下唯一。详见商户订单号
}

// 关闭订单的返回值
type CloseOrderResponse struct {
	ResponseModel
	// 当return_code为SUCCESS时
	ServiceResponseModel
	ResultMsg string `xml:"result_msg"` // 对业务结果的补充说明
}
