package check

import (
	"fmt"
	"net"
)

// Panic is
func Panic(e error) {
	if e != nil {
		panic(e)
	}
}

func TCPError(conn *net.TCPConn, e error) {
	if e != nil {
		conn.Write([]byte(e.Error()))
	}
}

func TCPErrorAndClose(conn *net.TCPConn, e error) {
	if e != nil {
		conn.Write([]byte(e.Error()))
		err := conn.Close()
		if err != nil {
			fmt.Println("TCP Error", err.Error())
		}
	}
}
