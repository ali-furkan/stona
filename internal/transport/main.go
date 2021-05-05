package transport

import (
	"net"

	"github.com/ali-furkqn/stona/internal/logger"
	"github.com/ali-furkqn/stona/internal/pkg/check"
)

type ServerConfig struct {
	Address string
	Port    int
	// TLS is optional. You can use it for TLS connection
	TLS struct {
		Cert string
		Key  string
	}
}

// Default Request Config Values
const (
	DefaultBodyLimit   = 32 * 1024
	DefaultConcurrency = 256 * 1024
)

func Serve(cfgServer *ServerConfig) {
	ln, err := net.ListenTCP("tcp", &net.TCPAddr{
		IP:   net.ParseIP(cfgServer.Address),
		Port: cfgServer.Port,
	})
	check.Panic(err)

	s := CreateServer()

	logger.Log("Transport", "Stona listening on ", cfgServer.Address, ":", cfgServer.Port)

	s.Serve(ln)
}
