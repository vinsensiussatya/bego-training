package context

type CtxKey string

const (
	Token     CtxKey = "token"
	RequestID CtxKey = "request_id"
)

func (c CtxKey) String() string {
	return string(c)
}
