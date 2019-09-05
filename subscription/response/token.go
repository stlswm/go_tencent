package response

type Token struct {
	Pub
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}
