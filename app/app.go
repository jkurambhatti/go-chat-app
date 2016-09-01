package app

import (
	"github.com/astaxie/beego/session"
)

var (
	GlobalSessions *session.Manager
)

func Init() {
	GlobalSessions, _ = session.NewManager("memory", `{"cookieName":"gosessionid","gclifetime":3600}`)
	go GlobalSessions.GC()
}
