package response

import "github.com/stlswm/go_tencent/subscription/response"

type JsApiTicket struct {
	response.Pub
	Ticket    string `json:"ticket"`
	ExpiresIn int64  `json:"expires_in"`
}
