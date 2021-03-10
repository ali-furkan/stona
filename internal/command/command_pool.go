package command

import (
	"sync"
)

var cmds = sync.Map{}
