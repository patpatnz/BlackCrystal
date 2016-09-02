package hosts

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

var (
	ErrLineTooLong    = errors.New("Encountered a line that is too long while parsing the hosts file")
	ErrInvalidVarPair = errors.New("Invalid var pair, should be x=y")
	ErrHostFormat     = errors.New("Host line format error")
)

type hostsFile struct {
	sync.Mutex
	Hosts  map[string]*Host
	Groups map[string]*Group
}

func newHostFile() *hostsFile {
	return &hostsFile{
		Hosts:  make(map[string]*Host),
		Groups: make(map[string]*Group),
	}
}

func (h *hostsFile) newHost(hostname string, group *Group) *Host {
	host := &Host{
		Name:   hostname,
		Groups: make(map[string]*Group),
		Vars:   make(map[string]interface{}),
	}

	h.Lock()
	defer h.Unlock()

	h.Hosts[hostname] = host

	return host
}

func (h *hostsFile) newGroup(name string) *Group {
	g := &Group{
		Name:  name,
		Hosts: make(map[string]*Host),
		Vars:  make(map[string]interface{}),
	}

	h.Lock()
	defer h.Unlock()

	h.Groups[name] = g

	return g
}

// NewFromFile loads a traditional ansible hosts file
func NewFromFile(filename string) (Hosts, error) {
	hf := newHostFile()

	in, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	lines := bufio.NewReader(in)

	var myGroup *Group
	var inVars = false

	for {
		bline, isPrefix, err := lines.ReadLine()
		if isPrefix {
			log.Printf("Encountered a really long line, blark")
			return nil, ErrLineTooLong
		}
		if err != nil && err != io.EOF {
			return nil, err
		}
		line := strings.Trim(string(bline), " \t")
		if len(line) != 0 && line[0] == '[' {
			// find group or create new
			tmp := strings.Split(line[1:strings.Index(line, "]")], ":")
			inVars = false
			groupName := tmp[0]
			if len(tmp) == 2 {
				if tmp[1] == "vars" {
					inVars = true
				}
			}
			ok := false
			if _, ok = hf.Groups[groupName]; !ok {
				myGroup = hf.newGroup(groupName)
			}
			continue
		} else if len(line) != 0 && line[0] != '#' {
			if inVars {
				// process vars in format x=y
				s := strings.Split(line, "=")
				if len(s) != 2 {
					return nil, ErrInvalidVarPair
				}
				k := strings.Trim(s[0], " ")
				v := strings.Trim(s[1], " ")
				myGroup.Vars[k] = v
			} else {
				fields := strings.Fields(line)
				if len(fields) < 1 {
					return nil, ErrHostFormat
				}
				if host, ok := hf.Hosts[fields[0]]; ok {
					host.setVars(fields[1:])
					if myGroup != nil {
						host.addToGroup(myGroup)
					}
				} else {
					host := hf.newHost(fields[0], myGroup)
					host.setVars(fields[1:])
				}
			}
		}
		if err != nil && err == io.EOF {
			return hf, nil
		}
	}
}

func (h *hostsFile) GetHosts(keyword string) []*Host {
	h.Lock()
	defer h.Unlock()
	if h, ok := h.Hosts[keyword]; ok {
		r := make([]*Host, 1)
		r[0] = h
		return r
	}
	if _, ok := h.Groups[keyword]; ok {
		b := make([]*Host, len(h.Hosts))
		i := 0
		for x := range h.Hosts {
			b[i] = h.Hosts[x]
		}
		return b
	}
	return make([]*Host, 0)
}
