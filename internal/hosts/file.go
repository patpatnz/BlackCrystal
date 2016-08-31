package hosts

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"strings"
)

var (
	ErrLineTooLong = errors.New("Encountered a line that is too long while parsing the hosts file")
)

type hostsFile struct {
	Hosts  map[string]*Host
	Groups map[string]*Group
}

func newHostFile() *hostsFile {
	return &hostsFile{
		Hosts:  make(map[string]*Host),
		Groups: make(map[string]*Group),
	}
}

func NewFromFile(filename string) (Hosts, error) {
	hf := newHostFile()

	in, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	lines := bufio.NewReader(in)

	var myGroup *Group = nil
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
				hf.Groups[groupName] = &Group{}
			}
			myGroup = hf.Groups[groupName]
			continue
		} else {
			log.Printf("line")
			_ = myGroup
			_ = inVars
		}
		if err != nil && err == io.EOF {
			return hf, nil
		}
	}
}

func (h hostsFile) GetHosts(keyword string) ([]*Host, error) {
	if h, ok := h.Hosts[keyword]; ok {
		r := make([]*Host, 1)
		r[0] = h
		return r, nil
	}
	if g, ok := h.Groups[keyword]; ok {
		return g.Hosts, nil
	}
	return nil, nil
}
