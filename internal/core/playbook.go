package core

type Playbook struct {
}

func NewPlaybookFromFile(filename string) (*Playbook, error) {
	playbook := &Playbook{}

	return playbook, nil
}
