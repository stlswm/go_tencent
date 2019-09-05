package account

import (
	"github.com/stlswm/go_tencent/subscription/account"
	"testing"
)

func TestGetToken(t *testing.T) {
	err, token := account.GetToken("", "")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(token)
	}
}
