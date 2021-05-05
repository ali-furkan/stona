package command

import (
	"fmt"

	"github.com/ali-furkqn/stona/internal/transport/message"
)

func Handle(msg *message.Message) *message.Message {

	cmd, exist := cmds.Load(msg.Command)
	if !exist {
		return message.NewResponseMessage(message.InvalidCommand, "Commands", "Command is invalid", nil, nil)
	}

	res, err := cmd.(BaseCommand).Run(msg)

	if err != nil {
		return message.NewResponseMessage(message.Error, fmt.Sprintf("%s Command", msg.Command), err.Error(), nil, nil)
	}

	return res
}
