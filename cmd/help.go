package cmd

import (
	"fmt"
	"strings"

	"github.com/ali-furkqn/stona/internal/cli"
)

type HelpCommand struct{}

func init() {
	_, err := Commands.InitCommand(&cli.CommandConfig{
		CommandName: "help",
		Cmd:         HelpCommand{},
	})
	if err != nil {
		panic(fmt.Sprintf("Error: %s", err))
	}
}

func (t HelpCommand) Run(args []string) error {

	if len(args) > 1 && Commands.Commands[args[0]] != nil {

		cmd := Commands.Commands[args[0]]

		fmt.Println(cmd.Help())

		return nil
	}

	text := `
Stona CLI

It allows storage initialization, buckets adding or deleting, 
configuration, logging, and more with just one command

Usage: stona <command> [args]

Commands:
`
	var cmds []string

	for name, cmd := range Commands.Commands {
		if name == "help" {
			continue
		}
		cmds = append(cmds, fmt.Sprintf("   %s\t%s", name, cmd.Synopsis()))
	}

	fmt.Printf("%s\n%s\n", strings.TrimSpace(text), strings.Join(cmds, "\n"))
	return nil
}

func (t HelpCommand) Synopsis() string {
	return "Prints Stona CLI version"
}

func (t HelpCommand) Help() string {
	text := `
Usage: stona help [arg]

Help provides help for any command in application.
You can type stona help [path to command] for full details

Flags:
    -h, --help help for help
`
	return strings.TrimSpace(text)
}
