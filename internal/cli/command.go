package cli

// Commands is Commands of Context for CLI
type Commands struct {
	Commands map[string]BaseCLICommand
}

// BaseCommand must be provided to Command in order to be runnable at CLI
type BaseCLICommand interface {
	// Help Function provides to explains what the command is with a long description
	Help() string

	// Execution Function of Command
	Run(args []string) error

	// Synopsis should return one-line, short synopsis of command
	Synopsis() string
}

// CommandConfig is Initialization Config
type CommandConfig struct {
	CommandName string
	// You can add with Aliases if you need to multiple tag name
	Aliases []string
	// Command Methods
	Cmd BaseCLICommand
}

// InitCommand inits new Command as your configure it in the Commands pool.
func (c *Commands) InitCommand(config *CommandConfig) (BaseCLICommand, error) {

	c.Commands[config.CommandName] = config.Cmd

	if config.Aliases != nil {
		for _, alias := range config.Aliases {
			c.Commands[alias] = config.Cmd
		}
	}

	return config.Cmd, nil
}
