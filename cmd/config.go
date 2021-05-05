package cmd

import (
	"fmt"
	"strings"

	"github.com/ali-furkqn/stona/internal/cli"
)

type ConfigCommand struct{}

func init() {
	_, err := Commands.InitCommand(&cli.CommandConfig{
		CommandName: "config",
		Cmd:         ConfigCommand{},
	})
	if err != nil {
		panic(fmt.Sprintf("Error: %s", err))
	}
}

func (t ConfigCommand) Run(args []string) error {
	fmt.Printf("Stona CLI %s", "Config Command")
	return nil
}

func (t ConfigCommand) Synopsis() string {
	return "Configure a running Stona server"
}

func (t ConfigCommand) Help() string {
	text := `
Usage: stona config

There are no arguments or flags to command.
`
	return strings.TrimSpace(text)
}
