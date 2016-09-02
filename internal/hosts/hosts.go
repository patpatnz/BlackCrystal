package hosts

import (
	"strings"
	"sync"
)

type Hosts interface {
	GetHosts(keyword string) []*Host
}

type Host struct {
	sync.RWMutex

	Name   string
	Groups map[string]*Group
	Vars   map[string]interface{}
}

type Group struct {
	sync.RWMutex

	Name  string
	Hosts map[string]*Host
	Vars  map[string]interface{}
}

func (h *Host) setVars(vars []string) error {
	h.Lock()
	defer h.Unlock()

	for x := range vars {
		s := strings.Split(vars[x], "=")
		if len(s) != 2 {
			return ErrInvalidVarPair
		}
		k := strings.Trim(s[0], " ")
		v := strings.Trim(s[1], " ")
		h.Vars[k] = v
	}
	return nil
}

func (h *Host) addToGroup(g *Group) {
	h.Lock()
	defer h.Unlock()
	g.Lock()
	defer g.Unlock()

	g.Hosts[h.Name] = h
	h.Groups[g.Name] = g
}

func (g *Group) addHost(h *Host) {
	h.Lock()
	defer h.Unlock()
	g.Lock()
	defer g.Unlock()

	g.Hosts[h.Name] = h
	h.Groups[g.Name] = g
}
