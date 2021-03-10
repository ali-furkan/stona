package cli

import (
	"flag"
	"fmt"
	"strings"
)

// handle method is where all commands are handled
func handle(args []string, commands map[string]BaseCLICommand) int {
	stdIn := strings.Join(args, " ")

	cmdFlag := flag.NewFlagSet("", flag.ExitOnError)
	hasHelpFlag := cmdFlag.Bool("help", false, "Help message for command")
	hasHelpAFlag := cmdFlag.Bool("h", false, "Help message for command")

	for path, command := range commands {
		if strings.HasPrefix(stdIn, path) {
			cmdFlag.Parse(args[1:])

			if *hasHelpFlag || *hasHelpAFlag {
				fmt.Println(command.Help())
				return 0
			}

			err := command.Run(args[1:])
			if err != nil {
				fmt.Printf("Error: %s", err)
				return 1
			}
			return 0
		}
	}

	err := commands["help"].Run(args)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return 1
	}

	return 0
}

// Init Method initializes CLI Application
func Init(args []string, commands map[string]BaseCLICommand) int {
	return handle(args, commands)
}
