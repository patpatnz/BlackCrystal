package lineinfile

import "github.com/patpatnz/BlackCrystal/internal/core"

func init() {
	core.CommandRegister(cmdLineInFile)
}

// This is the description of the core command
var cmdLineInFile = core.Command{
	Name: "LineInFile",
	//Parameters: {},
	GetInstance: func() (core.CommandIntf, error) {
		return &LineInFileCommand{}, nil
	},
}

type LineInFileCommand struct {
}

func (uc *LineInFileCommand) Setup(params []core.CommandParameter) error {
	return nil
}
