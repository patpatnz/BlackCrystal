package shell

import "github.com/patpatnz/BlackCrystal/internal/core"

func init() {
	core.CommandRegister(cmdShell)
}

// This is the description of the core command
var cmdShell = core.Command{
	Name: "Shell",
	//Parameters: {},
	GetInstance: func() (core.CommandIntf, error) {
		return &ShellCommand{}, nil
	},
}

type ShellCommand struct {
}

func (uc *ShellCommand) Setup(params []core.CommandParameter) error {
	return nil
}
