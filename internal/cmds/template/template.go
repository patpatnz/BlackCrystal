package template

import "github.com/patpatnz/BlackCrystal/internal/core"

func init() {
	core.CommandRegister(cmdTemplate)
}

// This is the description of the core command
var cmdTemplate = core.Command{
	Name: "Template",
	//Parameters: {},
	GetInstance: func() (core.CommandIntf, error) {
		return &TemplateCommand{}, nil
	},
}

type TemplateCommand struct {
}

func (uc *TemplateCommand) Setup(params []core.CommandParameter) error {
	return nil
}
