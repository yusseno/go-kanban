package session

var SessionStore Store

type SessionToken struct {
	TokenString string
}

type Store interface {
	Get(string) (SessionToken, error)
	Set(string, SessionToken) error
	Del(string) error
}
