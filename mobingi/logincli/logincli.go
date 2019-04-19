package logincli

import "github.com/mobingi/gosdk/mobingi/session"

// Context is our login abstraction on top of Session for cli-based operations.
type Context struct {
	session *session.Session
}
