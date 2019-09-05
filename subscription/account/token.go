package account

import (
	"encoding/json"
	"errors"
	"github.com/stlswm/go_tencent/subscription/account/response"
	"github.com/stlswm/go_tencent/tool"
)

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
