package message

import (
	"regexp"
	"strings"
)

func ParseMessage(cmdLine string) *Message {
	parts := strings.Split(cmdLine, " ")
	if len(parts) < 1 {
		return nil
	}

	if parts[0] != "STONA" {
		return nil
	}

	msg := &Message{}

	msg.Command = strings.TrimSpace(parts[1])

	if len(parts) == 2 || strings.TrimSpace(strings.Join(parts[2:], "")) == "" {
		return msg
	}

	// If message has header, parses it
	if match, _ := regexp.MatchString("(?:[\x5B])(.+)(?:[\x5D])", parts[2]); match {
		if len(parts) > 2 {
			msg.Params = make([][]byte, 16)
			for n, val := range strings.Split(strings.TrimSpace(strings.Join(parts[3:], "")), ",") {
				msg.Params[n] = []byte(val)
			}
		}

		msg.Header = make(map[string]string)
		fields := strings.Split(strings.Trim(parts[2], "[]"), ",")

		for _, field := range fields {
			fieldParams := strings.Split(field, "=")
			if len(fieldParams) > 1 {
				msg.Header[fieldParams[0]] = fieldParams[1]
			}
		}

	} else {
		msg.Params = make([][]byte, 16)

		ps := strings.Split(strings.TrimSpace(strings.Join(parts[2:], "")), ",")

		for n, val := range ps {
			if val != "" {
				msg.Params[n] = []byte(val)
			}
		}
	}

	return msg
}
