package cmd

import (
	"fmt"
	"strings"

	"github.com/ali-furkqn/stona/internal/cli"
	"github.com/ali-furkqn/stona/internal/version"
)

type versionCommand struct{}

func init() {
	_, err := Commands.InitCommand(&cli.CommandConfig{
		CommandName: "version",
		Cmd:         versionCommand{},
	})
	if err != nil {
		panic(fmt.Sprintf("Error: %s", err))
	}
}

func (t versionCommand) Run(args []string) error {
	fmt.Printf("Stona CLI %s", version.GetVersion())
	return nil
}

func (t versionCommand) Synopsis() string {
	return "Prints Stona CLI version"
}

func (t versionCommand) Help() string {
	text := `
Usage: stona version

  Prints version of Stona CLI.

  Prints the version:

    $ stona version

There are no arguments or flags to command.
`
	return strings.TrimSpace(text)
}
