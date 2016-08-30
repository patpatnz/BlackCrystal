package pip

import "github.com/patpatnz/BlackCrystal/internal/core"

func init() {
	core.CommandRegister(cmdPip)
}

// This is the description of the core command
var cmdPip = core.Command{
	Name: "Pip",
	//Parameters: {},
	GetInstance: func() (core.CommandIntf, error) {
		return &PipCommand{}, nil
	},
}

type PipCommand struct {
}

func (uc *PipCommand) Setup(params []core.CommandParameter) error {
	return nil
}
