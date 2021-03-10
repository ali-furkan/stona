package transport

import (
	"math"
	"sync"
)

type ConnectionPool struct {
	mux sync.Mutex

	open        uint32
	connections map[uint32]*Connection
}

func NewConnPool() *ConnectionPool {
	return &ConnectionPool{
		connections: make(map[uint32]*Connection),
	}
}

func (s *ConnectionPool) Add(conn *Connection) (id uint32, connection *Connection) {
	defer s.mux.Unlock()

	s.mux.Lock()
	{
		s.open++

		s.connections[s.open] = conn
		s.connections[s.open].id = s.open

		// This statement will removed when stona has role-based operations in the future
		s.connections[s.open].Perm = math.MaxUint16

		if s.connections[s.open].Path == "" {
			s.connections[s.open].Path = "/"
		}
	}

	return s.open, s.connections[s.open]
}

func (s *ConnectionPool) Delete(id uint32) {
	defer s.mux.Unlock()

	s.mux.Lock()
	conn := s.connections[id]

	if conn != nil {
		delete(s.connections, id)

		s.open--
	}
	return
}

func (s *ConnectionPool) GetCurrSize() uint32 {
	defer s.mux.Unlock()

	s.mux.Lock()
	open := s.open
	return open
}

func (s *ConnectionPool) GetAll() map[uint32]*Connection {
	defer s.mux.Unlock()

	s.mux.Lock()
	conns := s.connections

	return conns
}

func (s *ConnectionPool) Get(id uint32) *Connection {
	defer s.mux.Unlock()

	var conn *Connection
	s.mux.Lock()
	{
		if s.connections[id] != nil {
			conn = s.connections[id]
		}
	}
	return conn
}
