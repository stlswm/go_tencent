package account

import (
	"github.com/stlswm/go_tencent/subscription/account"
	"testing"
)

func TestJsApiTicket(t *testing.T) {
	token := "25_Ttg0nHIvK1mDEC9hQb7oMfVL80UsM0OFhfE0sLT-iUjTAty1z6kOtQ8qidNZ158Kr"
	token += "_Tb7EXe42W8P10Z7t8T-z9lqcLDwE19JAujbnTSNlQP8WCPiXRp39vG8vduqwJBohsaIZtGQs_2hWuyKADhAGATJM"
	err, ticket := account.JsApiTicket(token)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(ticket)
	}
}
