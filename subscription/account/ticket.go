package account

import (
	"encoding/json"
	"errors"
	"github.com/stlswm/go_tencent/subscription/account/response"
	"github.com/stlswm/go_tencent/tool"
)

// 获取js api ticket
func JsApiTicket(accessToken string) (error, *response.JsApiTicket) {
	url := "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=" + accessToken + "&type=jsapi"
	err, res := tool.HttpGet(url, nil)
	if err != nil {
		return err, nil
	}
	resTicket := &response.JsApiTicket{}
	err = json.Unmarshal(res, resTicket)
	if err != nil {
		return err, nil
	}
	if resTicket.ErrCode != 0 {
		return errors.New(resTicket.ErrMsg), nil
	}
	return nil, resTicket
}
