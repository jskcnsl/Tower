package server

import (
	"local/proxy_server/handler"
	"log"
	"net/http"
)

type Server struct {
	ListenAddress string
	SocketServer  bool
	TLS           bool
	Handler       handler.Handler
}

func (server *Server) EasyServer() error {
	defer log.Printf("[x] server gg\n")
	http.HandleFunc("/", server.Handler.Handle)
	if server.TLS {
		// TODO: add cert file content
		if err := http.ListenAndServeTLS(server.ListenAddress, "", "", nil); err != nil {
			log.Fatal("[x] ListenAndServe: ", err)
			return err
		}
	} else {
		log.Printf("[*] easy server start listen %s\n", server.ListenAddress)
		if err := http.ListenAndServe(server.ListenAddress, nil); err != nil {
			log.Fatal("[x] ListenAndServe: ", err)
			return err
		}
	}
	return nil
}

func NewServer(listenAddress string, tls bool, reqHandler handler.Handler) *Server {
	return &Server{
		ListenAddress: listenAddress,
		TLS:           tls,
		Handler:       reqHandler,
		SocketServer:  false,
	}
}
