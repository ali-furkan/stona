package cmd

import "github.com/ali-furkqn/stona/internal/cli"

// Commands is a Global Commands Context
var Commands = &cli.Commands{
	Commands: make(map[string]cli.BaseCLICommand),
}

func Run(args []string) int {
	return cli.Init(args, Commands.Commands)
}
