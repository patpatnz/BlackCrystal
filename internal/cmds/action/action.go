package action

import "github.com/patpatnz/BlackCrystal/internal/core"

func init() {
	core.CommandRegister(cmdAction)
}

// This is the description of the core command
var cmdAction = core.Command{
	Name: "Action",
	//Parameters: {},
	GetInstance: func() (core.CommandIntf, error) {
		return &ActionCommand{}, nil
	},
}

type ActionCommand struct {
}

func (uc *ActionCommand) Setup(params []core.CommandParameter) error {
	return nil
}
