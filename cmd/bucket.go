package cmd

import (
	"fmt"
	"strings"

	"github.com/ali-furkqn/stona/internal/cli"
)

type BucketCommand struct{}

func init() {
	_, err := Commands.InitCommand(&cli.CommandConfig{
		CommandName: "bucket",
		Cmd:         BucketCommand{},
	})
	if err != nil {
		panic(fmt.Sprintf("Error: %s", err))
	}
}

func (t BucketCommand) Run(args []string) error {
	fmt.Printf("Stona CLI %s", "Bucket Command")
	return nil
}

func (t BucketCommand) Synopsis() string {
	return "Manipulate the Buckets."
}

func (t BucketCommand) Help() string {
	text := `
Usage: stona bucket <command> [args]

	Manipulates the Bucket with CRUD actions 
	when you connected any stona server.

	Create Bucket:

		$ stona bucket create <name> --config <config-file>

	Delete Bucket:
		
		$ stona bucket delete <name>

	Get Bucket:

		$ stona bucket get [name] --list

	
	Update Bucket:

		$ stona bucket update

	Flags:
		-c, --config: Loads config file for Bucket
		-l, --list: Gets all of the bucket from the Stona Server
`
	return strings.TrimSpace(text)
}
