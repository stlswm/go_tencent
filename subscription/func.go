package subscription

import (
	"encoding/json"
	"errors"
	"github.com/stlswm/go_tencent/subscription/response"
	"github.com/stlswm/go_tencent/tool"
)

type TokenResponse struct {
}

// 获取token
func GetToken(appId string, secret string) (error, *response.Token) {
	url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + appId + "&secret=" + secret
	err, res := tool.HttpGet(url, nil)
	if err != nil {
		return err, nil
	}
	resToken := &response.Token{}
	err = json.Unmarshal(res, resToken)
	if err != nil {
		return err, nil
	}
	if resToken.ErrCode != 0 {
		return errors.New(resToken.ErrMsg), nil
	}
	return nil, resToken
}
