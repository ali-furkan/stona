package cmd

import (
	"fmt"
	"strings"

	"github.com/ali-furkqn/stona/internal/cli"
	"github.com/ali-furkqn/stona/internal/version"
)

type ConnectCommand struct{}

func init() {
	_, err := Commands.InitCommand(&cli.CommandConfig{
		CommandName: "connect",
		Cmd:         ConnectCommand{},
	})
	if err != nil {
		panic(fmt.Sprintf("Error: %s", err))
	}
}

func (t ConnectCommand) Run(args []string) error {
	fmt.Printf("Stona CLI %s", version.GetVersion())
	return nil
}

func (t ConnectCommand) Synopsis() string {
	return "Connect to any Stona Server"
}

func (t ConnectCommand) Help() string {
	text := `
Usage: stona connect <address>

Connect to any Stona Server

Example:

	$ stona connect my-awesome-storage.com

There are no arguments or flags to command.
`
	return strings.TrimSpace(text)
}
