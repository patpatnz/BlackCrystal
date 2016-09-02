package transport

type docker struct {
}

func init() {
	Register(&docker{})
}

func (d docker) GetName() string {
	return "Docker"
}

func (d docker) GetInstance() (Transport, error) {
	return &docker{}, nil
}
