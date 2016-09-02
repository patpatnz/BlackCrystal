package transport

type ssh struct {
}

func init() {
	Register(&ssh{})
}

func (s ssh) GetName() string {
	return "SSH"
}

func (s ssh) GetInstance() (Transport, error) {
	return &ssh{}, nil
}
