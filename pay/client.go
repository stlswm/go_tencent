package pay

// 支付客户端
type Client struct {
	mchId string
	appId string
}

// 获取商户号
func (c *Client) GetMchId() string {
	return c.mchId
}

// 获取应用id
func (c *Client) GetAppId() string {
	return c.appId
}

func NewClient() *Client {
	c := &Client{}
	return c
}
