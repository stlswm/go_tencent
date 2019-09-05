package subscription

import (
	"github.com/stlswm/go_tencent/subscription"
	"testing"
)

func TestGetToken(t *testing.T) {
	err, token := subscription.GetToken("", "")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(token)
	}
}
