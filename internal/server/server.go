package server

import (
	"net"
	"net/http"

	"github.com/sail-host/cloud/internal/global"
	"github.com/sail-host/cloud/internal/init/router"
)

func Start(devMode bool) {
	Init(devMode)

	appRouter := router.Routers()

	server := &http.Server{
		Addr:    global.CONF.System.BindAddress + ":" + global.CONF.System.Port,
		Handler: appRouter,
	}
	ln, err := net.Listen("tcp", server.Addr)
	if err != nil {
		panic(err)
	}

	// TODO: add ssl config setup
	type tcpKeepAliveListener struct {
		*net.TCPListener
	}
	global.LOG.Info("Server is running on ", server.Addr)
	if err := server.Serve(tcpKeepAliveListener{ln.(*net.TCPListener)}); err != nil {
		panic(err)
	}
}
