package signature

type SInterface interface {
	GetMethod() string
	GetVersion() string
	GetType() string
	Sign(data string, secret string) string
}
