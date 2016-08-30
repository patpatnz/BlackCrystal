package docker

import "github.com/patpatnz/BlackCrystal/internal/core"

func init() {
	core.CommandRegister(cmdDocker)
}

// This is the description of the core command
var cmdDocker = core.Command{
	Name: "Docker",
	//Parameters: {},
	GetInstance: func() (core.CommandIntf, error) {
		return &DockerCommand{}, nil
	},
}

type DockerCommand struct {
}

func (uc *DockerCommand) Setup(params []core.CommandParameter) error {
	return nil
}
