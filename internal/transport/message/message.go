package message

import (
	"fmt"
	"strconv"
	"strings"
)

// Example Message: STONA COMMAND [PATH=/] p0,p1,p2,pn

type Message struct {
	Command string
	Header  map[string]string
	Params  [][]byte
	Status  MessageStatus
	Context string
	// Type of Message is Response, If Type is true.
	// If It's false, type is request
	Type bool
}

func (m *Message) Byte() []byte {
	return []byte(m.String())
}

func (m *Message) String() string {
	msg := "STONA"

	params := ""
	for _, val := range m.Params {
		if val != nil {
			params += fmt.Sprintf("%s,", string(val))
		}
	}

	// If Type of Message is Response, Adds Status Section
	if m.Type {
		status := strconv.FormatUint(uint64(m.Status), 10)
		msg += " " + status
		if m.Context != "" {
			msg += " " + m.Context
		}
	} else {
		msg += " " + m.Command
	}

	if m.Header != nil {
		header := ""
		i := 0
		for key, val := range m.Header {
			i++
			header += fmt.Sprintf("%s=%s", key, val)
			if len(m.Header) < i {
				header += ","
			}
		}

		msg += fmt.Sprintf(" [%s]", header)
	}

	msg += " " + params

	return strings.TrimSpace(msg)
}

func NewResponseMessage(status MessageStatus, ctx string, desc string, defHead map[string]string, data []byte) *Message {
	var header map[string]string

	if defHead != nil {
		header = defHead
	}

	if desc != "" {
		if defHead == nil {
			header = make(map[string]string)
		}
		header["Description"] = desc
	}

	return &Message{
		Context: ctx,
		Status:  status,
		Params:  [][]byte{data},
		Header:  header,
		Type:    true,
	}
}

func NewRequestMessage(cmd string, defHead map[string]string, params [][]byte) *Message {
	var header map[string]string

	if defHead != nil {
		header = defHead
	}

	return &Message{
		Command: cmd,
		Params:  params,
		Header:  header,
	}
}
