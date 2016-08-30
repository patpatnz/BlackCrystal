package assert

import "github.com/patpatnz/BlackCrystal/internal/core"

func init() {
	core.CommandRegister(cmdAssert)
}

// This is the description of the core command
var cmdAssert = core.Command{
	Name: "Assert",
	//Parameters: {},
	GetInstance: func() (core.CommandIntf, error) {
		return &AssertCommand{}, nil
	},
}

type AssertCommand struct {
}

func (uc *AssertCommand) Setup(params []core.CommandParameter) error {
	return nil
}
