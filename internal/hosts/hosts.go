package hosts

type Hosts interface {
	GetHosts(keyword string) ([]*Host, error)
}

type Host struct {
	Groups []*Group
	Vars   map[string]interface{}
}

type Group struct {
	Hosts []*Host
	Vars  map[string]interface{}
}
