package user

import "github.com/patpatnz/BlackCrystal/internal/core"

func init() {
	core.CommandRegister(cmdUser)
}

// This is the description of the core command
var cmdUser = core.Command{
	Name: "User",
	//Parameters: {},
	GetInstance: func() (core.CommandIntf, error) {
		return &UserCommand{}, nil
	},
}

type UserCommand struct {
}

func (uc *UserCommand) Setup(params []core.CommandParameter) error {
	return nil
}
