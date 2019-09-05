package response

import "github.com/stlswm/go_tencent/subscription/response"

type Token struct {
	response.Pub
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}
