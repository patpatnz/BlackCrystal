package file

import "github.com/patpatnz/BlackCrystal/internal/core"

func init() {
	core.CommandRegister(cmdFile)
}

// This is the description of the core command
var cmdFile = core.Command{
	Name: "File",
	//Parameters: {},
	GetInstance: func() (core.CommandIntf, error) {
		return &FileCommand{}, nil
	},
}

type FileCommand struct {
}

func (uc *FileCommand) Setup(params []core.CommandParameter) error {
	return nil
}
