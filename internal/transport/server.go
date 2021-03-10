package transport

import (
	"net"
	"sync"

	"golang.org/x/time/rate"
)

type Server struct {
	ln net.Listener

	openConn int32
	connPool *ConnectionPool

	visitorLimits map[string]*rate.Limiter

	BodyLimit      int
	MaxConcurrency int

	mux sync.Mutex
}

func CreateServer() *Server {
	connPool := NewConnPool()

	return &Server{
		BodyLimit:      DefaultBodyLimit,
		MaxConcurrency: DefaultConcurrency,
		connPool:       connPool,
		visitorLimits:  make(map[string]*rate.Limiter),
	}
}
