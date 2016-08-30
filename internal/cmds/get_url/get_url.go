package user

import "github.com/patpatnz/BlackCrystal/internal/core"

func init() {
	core.CommandRegister(cmdGetUrl)
}

// This is the description of the core command
var cmdGetUrl = core.Command{
	Name: "Get_Url",
	//Parameters: {},
	GetInstance: func() (core.CommandIntf, error) {
		return &GetUrlCommand{}, nil
	},
}

type GetUrlCommand struct {
}

func (uc *GetUrlCommand) Setup(params []core.CommandParameter) error {
	return nil
}
