package apt

import "github.com/patpatnz/BlackCrystal/internal/core"

func init() {
	core.CommandRegister(cmdApt)
}

// This is the description of the core command
var cmdApt = core.Command{
	Name: "Apt",
	//Parameters: {},
	GetInstance: func() (core.CommandIntf, error) {
		return &AptCommand{}, nil
	},
}

type AptCommand struct {
}

func (uc *AptCommand) Setup(params []core.CommandParameter) error {
	return nil
}
