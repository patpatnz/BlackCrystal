package service

import "github.com/patpatnz/BlackCrystal/internal/core"

func init() {
	core.CommandRegister(cmdService)
}

// This is the description of the core command
var cmdService = core.Command{
	Name: "Service",
	//Parameters: {},
	GetInstance: func() (core.CommandIntf, error) {
		return &ServiceCommand{}, nil
	},
}

type ServiceCommand struct {
}

func (uc *ServiceCommand) Setup(params []core.CommandParameter) error {
	return nil
}
