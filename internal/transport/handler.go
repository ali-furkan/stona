package transport

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/ali-furkqn/stona/internal/command"
	"github.com/ali-furkqn/stona/internal/config"
	"github.com/ali-furkqn/stona/internal/logger"
	"github.com/ali-furkqn/stona/internal/pkg/check"
	"github.com/ali-furkqn/stona/internal/transport/message"
	"golang.org/x/time/rate"
)

func (s *Server) aliveTimeout(conn *net.TCPConn, i int) func() {
	i = i + 1

	return func() {
		buf := make([]byte, s.BodyLimit)

		n, err := conn.Read(buf)

		if n == 0 || err != nil {
			return
		}

		if s.isSuccAuth(string(buf)) {
			return
		}

		if i > 2 {
			logger.Log("Connection", conn.RemoteAddr(), " Connection Timed out")
			if _, err := conn.Write(message.NewResponseMessage(message.ConnectionTimeout, "Connection", "Connection Timed out", nil, nil).Byte()); err != nil {
				return
			}
			conn.Close()
			return
		}

		msg := fmt.Sprintf("Attention! Stona is waiting for client to verify itself [REPEAT: %d ] \n", i)
		if _, err := conn.Write(message.NewResponseMessage(message.Warn, "Connection", msg, nil, nil).Byte()); err != nil {
			return
		}
		time.AfterFunc(time.Duration(config.Config().Connection.Timeout/3)*time.Second, s.aliveTimeout(conn, i))
	}
}

func (s *Server) isSuccAuth(data string) bool {
	if strings.Split(data, " ")[0] != "STONA" {
		return false
	}

	svUsername := config.Config().Authentication.Username
	svPassword := config.Config().Authentication.Password

	if svUsername == "" || svPassword == "" {
		return true
	}

	cmd := message.ParseMessage(strings.Join(strings.Split(data, " ")[1:], " "))
	if cmd == nil {
		return false
	}

	if cmd.Command != "AUTH" || len(cmd.Params) != 2 {
		return false
	}

	username := string(cmd.Params[0])
	password := string(cmd.Params[1])

	if username == svUsername && password == svPassword {
		return true
	}

	return false
}

func isHTTPRequest(data []byte) (isHTTP bool) {
	args := strings.Split(strings.Replace(string(data), "\r\n", "\n", -1), "\n")
	if len(args) < 2 {
		isHTTP = false
		return
	}

	versionOfProtcolParts := strings.Split(strings.Split(args[0], " ")[2], "/")
	if len(versionOfProtcolParts) < 2 {
		isHTTP = false
		return
	}

	protocol := strings.TrimSpace(versionOfProtcolParts[0])
	ver, err := strconv.ParseFloat(strings.TrimSpace(versionOfProtcolParts[1]), 16)

	if err != nil {
		isHTTP = false
		return
	}

	if strings.ToUpper(protocol) == "HTTP" {
		if ver == 1.1 || ver == 1.2 || ver == 2.0 {
			isHTTP = true
		}
		return
	}

	isHTTP = false
	return
}

// AddOrCheckVisitorLimit is check acceptability of the connection
func (s *Server) AddOrCheckVisitorLimit(conn *net.TCPConn) (err error) {
	defer s.mux.Unlock()

	s.mux.Lock()
	ip, _, err := net.SplitHostPort(conn.RemoteAddr().String())

	if err != nil {
		return err
	}

	limiter, isExist := s.visitorLimits[ip]

	if !isExist {
		rt := rate.Every(10 * time.Second)
		limiter = rate.NewLimiter(rt, 10)

		s.visitorLimits[ip] = limiter
	}

	if limiter.Allow() == false {
		return errors.New("Too many requests")
	}

	return nil
}

// HTTPHandler handles if the connection is http protocol
func (s *Server) HTTPHandler(conn *Connection) {
	defer func() {
		conn.Close()
		s.connPool.Delete(conn.id)
		return
	}()

	conn.Write(s.httpOk(1.2))
	return
}

// StonaHandler handles if the connection is stona protocol
func (s *Server) StonaHandler(conn *Connection) {
	defer func() {
		conn.Close()
		s.connPool.Delete(conn.id)
	}()

	if !conn.Write(message.NewResponseMessage(message.Successful, "Server", "Welcome to Stona", nil, nil).Byte()) {
		return
	}

	for {

		buf := conn.Read(s.BodyLimit)

		cmd := message.ParseMessage(string(buf))

		if cmd == nil {
			if !conn.Write(message.NewResponseMessage(message.InvalidCommand, "Handler", "Invalid Command", nil, nil).Byte()) {
				return
			}
			continue
		}

		res := command.Handle(cmd)

		if !conn.Write(res.Byte()) {
			break
		}
	}
}

// HandleConn handles connection as type (Http, Stona etc.)
// It also processes middlewares (rate-limit, authentication etc.)
// over the connection before handles connection
func (s *Server) HandleConn(conn *net.TCPConn) {
	buf := make([]byte, s.BodyLimit)

	n, err := conn.Read(buf)

	if n == 0 || err != nil {
		check.TCPErrorAndClose(conn, err)
		return
	}

	isHTTP := isHTTPRequest(buf)

	connection := &Connection{
		conn: conn,
	}

	err = s.AddOrCheckVisitorLimit(conn)

	if isHTTP {
		if err != nil {
			conn.Write(s.httpTooManyReq(1.2))
			conn.Close()
			return
		}

		s.HTTPHandler(connection)
		return
	}

	if err != nil {
		conn.Write(message.NewResponseMessage(message.TooManyRequest, "", "", nil, nil).Byte())
		conn.Close()
		return
	}

	time.AfterFunc(
		time.Duration(config.Config().Connection.Timeout/3)*time.Second,
		s.aliveTimeout(conn, 0),
	)

	if !s.isSuccAuth(string(buf)) {
		if _, err := conn.Write(message.NewResponseMessage(message.UnauthorizedConnection, "Authentication", "", nil, nil).Byte()); err != nil {
			return
		}
		conn.Close()
		return
	}

	s.connPool.Add(connection)

	s.StonaHandler(connection)
	return
}
