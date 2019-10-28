package signature

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

type Md5 struct {
}

func (s *Md5) GetMethod() string {
	return "MD5"
}

func (s *Md5) GetVersion() string {
	return "1.0"
}

func (s *Md5) GetType() string {
	return ""
}

func (s *Md5) Sign(data string, secret string) string {
	h := md5.New()
	h.Write([]byte(data))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}
