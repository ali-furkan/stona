package command

import (
	"fmt"

	"github.com/ali-furkqn/stona/internal/transport/message"
)

type pingCommand struct{}

func (cmd pingCommand) Name() string {
	return "PING"
}

func (cmd pingCommand) Run(msg *message.Message) (*message.Message, error) {
	return message.NewResponseMessage(message.Successful, "", "", nil, nil), nil
}

func init() {

	cmd := &pingCommand{}
	err := CreateCommand(cmd)

	if err != nil {
		panic(fmt.Sprintf("Error: %s", err))
	}
}
