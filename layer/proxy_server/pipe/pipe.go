package pipe

import (
	"local/proxy_server/handler/easy"
	serverPkg "local/proxy_server/server"
	"sync"
)

type Pipe struct {
	// ListenAddress is the address to be listen frontend call
	ListenAddress string
	// RequestAddress is the address to request after data transform by handler
	RequestAddress string
	// SocketServer will create server to handle transport layer
	SocketServer bool
	// TLS will create tls server
	TLS bool
}

func (pipe *Pipe) Start() error {
	var serverWg sync.WaitGroup
	handler := &easy.Handler{
		RequestAddress: pipe.RequestAddress,
	}
	server := serverPkg.NewServer(pipe.ListenAddress, pipe.TLS, handler)
	serverWg.Add(1)
	go func() {
		defer serverWg.Done()
		server.EasyServer()
	}()
	serverWg.Wait()
	return nil
}

func NewPipe(listenAddr, requestAddr string, socket, tls bool) *Pipe {
	return &Pipe{
		ListenAddress:  listenAddr,
		RequestAddress: requestAddr,
		SocketServer:   socket,
		TLS:            tls,
	}
}
