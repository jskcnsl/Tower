package proxy

import (
	"log"
	"sync"

	"github.com/jskcnsl/tower/proxy/handler/easy"
)

type Proxy struct {
	// ListenAddress is the address to be listen frontend call
	ListenAddress string
	// RequestAddress is the address to request after data transform by handler
	RequestAddress string
	// SocketServer will create server to handle transport layer
	SocketServer bool
	// TLS will create tls server
	TLS bool
}

// NewProxy create a new proxy
func NewProxy(listenAddr, requestAddr string, socket, tls bool) *Proxy {
	return &Proxy{
		ListenAddress:  listenAddr,
		RequestAddress: requestAddr,
		SocketServer:   socket,
		TLS:            tls,
	}
}

// Start run the proxy
func (p *Proxy) Start() {
	var proxyWG sync.WaitGroup

	// create goroutine to run server listen request,
	// and choose one handler to handle request.
	proxyWG.Add(1)
	go func() {
		// create easy handler to print log and send to server
		handler := &easy.Handler{
			RequestAddress: p.RequestAddress,
		}

		// create server to listen request
		server := NewServer(p.ListenAddress, p.TLS, handler)

		var serverWg sync.WaitGroup
		// run server in goroutine
		serverWg.Add(1)
		go func() {
			defer serverWg.Done()
			if err := server.EasyServer(); err != nil {
				log.Panicf("proxy server run failed: %s", err)
			}
		}()
		serverWg.Wait()
	}()
	proxyWG.Wait()
}
