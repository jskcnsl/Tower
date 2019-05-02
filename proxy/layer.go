package proxy

type Layer struct {
	ListenAddress  string
	RequestAddress string
	SocketServer   bool
	TLS            bool
	proxy          *Proxy
}

func (l *Layer) Config(args ...interface{}) error {
	l.ListenAddress = args[0].(string)
	l.RequestAddress = args[1].(string)
	l.SocketServer = args[2].(bool)
	l.TLS = args[3].(bool)
	return nil
}

func (l *Layer) Run(args ...interface{}) error {
	proxy := NewProxy(l.ListenAddress, l.RequestAddress, l.SocketServer, l.TLS)
	l.proxy = proxy
	return nil
}
