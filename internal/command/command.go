package command

import (
	"errors"

	"github.com/ali-furkqn/stona/internal/transport/message"
)

type BaseCommand interface {
	Name() string
	Run(msg *message.Message) (*message.Message, error)
}

func CreateCommand(cmd BaseCommand) error {
	if lcmd, _ := cmds.Load(cmd.Name()); lcmd != nil {
		return errors.New("Duplication Error. 'cmds' has more than one command with the same name ")
	}

	cmds.Store(cmd.Name(), cmd)

	return nil
}
