package request

// 发放红包请求体
type RedPacket struct {
	NonceStr    string `xml:"nonce_str"`
	Sign        string `xml:"sign"`
	MchBillNo   string `xml:"mch_billno"`
	MchId       string `xml:"mch_id"`
	WxAppId     string `xml:"wxappid"`
	SendName    string `xml:"send_name"`
	ReOpenId    string `xml:"re_openid"`
	TotalAmount int    `xml:"total_amount"`
	TotalNum    int    `xml:"total_num"`
	Wishing     string `xml:"wishing"`
	ClientIp    string `xml:"client_ip"`
	ActName     string `xml:"act_name"`
	Remark      string `xml:"remark"`
	SceneId     string `xml:"scene_id"`
	RiskInfo    string `xml:"risk_info"`
}
