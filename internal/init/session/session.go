package session

import (
	"github.com/sail-host/cloud/internal/global"
	"github.com/sail-host/cloud/internal/init/session/psession"
)

func Init() {
	global.SESSION = psession.NewPSession(global.CACHE)
	global.LOG.Info("init session successfully")
}
