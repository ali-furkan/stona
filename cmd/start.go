package cmd

import (
	"flag"
	"fmt"
	"strings"

	"github.com/ali-furkqn/stona/internal/cli"
	"github.com/ali-furkqn/stona/internal/runner"
)

type StartCommand struct{}

func init() {
	_, err := Commands.InitCommand(&cli.CommandConfig{
		CommandName: "start",
		Cmd:         StartCommand{},
	})
	if err != nil {
		panic(fmt.Sprintf("Error: %s", err))
	}
}

func (t StartCommand) Run(args []string) error {
	fmt.Printf("Stona CLI %s\n", "Start Command")
	cmdFlag := flag.NewFlagSet("", flag.ExitOnError)
	cfgPath := flag.String("config", "", "Configuration file to load")
	cmdFlag.Parse(args)

	runner.Start(*cfgPath)

	return nil
}

func (t StartCommand) Synopsis() string {
	return "Start the Stona Server with your configurations"
}

func (t StartCommand) Help() string {
	text := `
Usage: stona start -p <PORT> --config <CONFIG_FILE> --ui

Start the Stona server

   Example:

       $ stona start -p 8080 --config "./config.yml"


Flags:
    -p, --port:	Listen to port
    --config:	Config File
`
	return strings.TrimSpace(text)
}
