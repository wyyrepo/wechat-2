package wechat

import (
	"encoding/json"
)

// 向微信发送请求
func (c *Client) doWeChat(relativeUrl string, bodyObj interface{}) (bytes []byte, err error) {
	// 将bodyObj转换为map[string]interface{}类型
	bodyJson, _ := json.Marshal(bodyObj)
	body := make(map[string]interface{})
	_ = json.Unmarshal(bodyJson, &body)
	// 添加固定参数
	body["appid"] = c.config.AppId
	body["mch_id"] = c.config.MchId
	if c.isFacilitator() {
		body["sub_appid"] = c.config.SubAppId
		body["sub_mch_id"] = c.config.SubMchId
	}
	nonceStr := GetRandomString(32)
	body["nonce_str"] = nonceStr
	// 生成签名
	signType, _ := body["sign_type"].(string)
	var sign string
	if c.isProd {
		sign = c.localSign(body, signType, c.apiKey)
	} else {
		body["sign_type"] = SignTypeMD5
		key, iErr := c.sandboxSign(nonceStr, SignTypeMD5)
		if err = iErr; iErr != nil {
			return
		}
		sign = c.localSign(body, SignTypeMD5, key)
	}
	body["sign"] = sign
	// 发起请求
	bytes, err = httpPost(c.url(relativeUrl), GenerateXml(body))
	return
}
