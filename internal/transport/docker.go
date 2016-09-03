package transport

import "github.com/patpatnz/BlackCrystal/internal/hosts"

type docker struct {
}

func init() {
	//Register(&docker{})
}

func (d docker) GetName() string {
	return "Docker"
}

func (d docker) GetInstance(host *hosts.Host) (Transport, error) {
	//return &docker{}, nil
	return nil, nil
}

func (d docker) Run(cmd string) ([]string, error) {
	return nil, nil
}
