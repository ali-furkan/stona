package transport

import (
	"net"

	"github.com/ali-furkqn/stona/internal/logger"
)

// Serve is serve with TCP
func (s *Server) Serve(ln *net.TCPListener) {
	defer ln.Close()

	for {
		if s.connPool.open >= 5 {
			continue
		}

		conn, err := ln.AcceptTCP()
		if err != nil {
			logger.Log("Error", err)
			if err := conn.Close(); err != nil {
				logger.Log("Connection", "Failed to close listener: ", err)
			}
			continue
		}

		logger.Log("Connection", "Connected to ", conn.RemoteAddr())

		go s.HandleConn(conn)
	}
}
