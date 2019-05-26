package proxy

import (
	"github.com/jskcnsl/tower/command"
	"github.com/spf13/cobra"
)

var (
	listenAddr  string
	requestAddr string
	socket      bool
	tls         bool
)

// Layer define proxy layer struct
type Layer struct {
	ListenAddress  string
	RequestAddress string
	SocketServer   bool
	TLS            bool
	proxy          *Proxy
}

// Config set proxy layer
func (l *Layer) Config(args ...interface{}) error {
	l.ListenAddress = listenAddr
	l.RequestAddress = requestAddr
	l.SocketServer = socket
	l.TLS = tls
	return nil
}

// Run start proxy lyaer
func (l *Layer) Run(args ...interface{}) error {
	proxy := NewProxy(l.ListenAddress, l.RequestAddress, l.SocketServer, l.TLS)
	l.proxy = proxy
	l.proxy.Start()
	return nil
}

func init() {
	cmd := &cobra.Command{
		Use:   "proxy",
		Short: "Proxy layer start proxy server",
		Long: `Proxy layer start proxy server to listen on one port,
and transform request to the true handle server.`,
		Run: func(cmd *cobra.Command, args []string) {
			l := new(Layer)
			if err := l.Config(); err != nil {
				panic(err)
			}
			if err := l.Run(); err != nil {
				return
			}
		},
	}
	cmd.Flags().StringVar(&listenAddr, "listen", "0.0.0.0:8080", "proxy server listen on")
	cmd.Flags().StringVar(&requestAddr, "handle", "", "server to handle request")
	cmd.Flags().BoolVar(&socket, "socket", false, "whether handle request as socket")
	cmd.Flags().BoolVar(&tls, "tls", false, "whether run proxy server with tls")
	command.RegisterLayer(cmd)
}
