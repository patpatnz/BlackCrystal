package copy

import "github.com/patpatnz/BlackCrystal/internal/core"

func init() {
	core.CommandRegister(cmdCopy)
}

// This is the description of the core command
var cmdCopy = core.Command{
	Name: "Copy",
	//Parameters: {},
	GetInstance: func() (core.CommandIntf, error) {
		return &CopyCommand{}, nil
	},
}

type CopyCommand struct {
}

func (uc *CopyCommand) Setup(params []core.CommandParameter) error {
	return nil
}
