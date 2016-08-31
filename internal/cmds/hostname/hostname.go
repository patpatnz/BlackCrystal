package hostname

import "github.com/patpatnz/BlackCrystal/internal/core"

func init() {
	core.CommandRegister(cmdHostname)
}

// This is the description of the core command
var cmdHostname = core.Command{
	Name: "Hostname",
	//Parameters: {},
	GetInstance: func() (core.CommandIntf, error) {
		return &HostnameCommand{}, nil
	},
}

type HostnameCommand struct {
}

func (uc *HostnameCommand) Setup(params []core.CommandParameter) error {
	return nil
}
