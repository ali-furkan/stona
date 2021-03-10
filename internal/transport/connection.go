package transport

import (
	"net"

	"github.com/ali-furkqn/stona/internal/logger"
	"github.com/ali-furkqn/stona/internal/pkg/check"
)

type Connection struct {
	id   uint32
	Path string
	Perm uint16
	Auth bool
	conn *net.TCPConn
}

func (s *Connection) Read(bodyLimit int) []byte {
	buf := make([]byte, bodyLimit)

	n, err := s.conn.Read(buf)
	if n == 0 || err != nil {
		check.TCPError(s.conn, err)
		return nil
	}

	return buf
}

func (s *Connection) Write(buf []byte) bool {
	if _, err := s.conn.Write(buf); err != nil {
		logger.Log("Connection", "Failed to write: ", err)
		return false
	}

	return true
}

func (s *Connection) Close() {
	if s.conn == nil {
		return
	}

	err := s.conn.Close()

	if err != nil {
		logger.Log("Connection", "Closing Connnection: ", err)
	}
}
