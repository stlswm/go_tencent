package nonce

import (
	"github.com/stlswm/go_tencent/nonce"
	"log"
	"testing"
)

func TestStr(t *testing.T) {
	for i := 1; i <= 100; i++ {
		log.Println(nonce.Str(i))
	}
}
func TestNum(t *testing.T) {
	for i := 1; i <= 100; i++ {
		log.Println(nonce.Num(i))
	}
}
