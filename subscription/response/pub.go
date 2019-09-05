package response

// 公共响应头
type Pub struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
