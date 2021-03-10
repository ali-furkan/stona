package main

import (
	"os"

	"github.com/ali-furkqn/stona/cmd"
)

func main() {
	os.Exit(cmd.Run(os.Args[1:]))
}
