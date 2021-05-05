package transport

import (
	"fmt"

	"github.com/ali-furkqn/stona/internal/version"
)

func (s *Server) httpOk(httpVer float64) []byte {
	return []byte(fmt.Sprintf("HTTP/%g 200 OK\r\nContent-Type: text/plain\r\n\r\nStona version %s", httpVer, version.GetVersion()))
}

func (s *Server) httpBadReq(httpVer float64) []byte {
	return []byte(fmt.Sprintf("HTTP/%g 400 Bad Request\r\nContent-Type: text/plain\r\n\r\nBad Request", httpVer))
}

func (s *Server) httpTooManyReq(httpVer float64) []byte {
	return []byte(fmt.Sprintf("HTTP/%g 429 Too Many Requests\r\nContent-Type: text/plain\r\nRetry-After: 10\r\n\r\nToo Many Request", httpVer))
}
