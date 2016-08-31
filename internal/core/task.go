package core

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/kylelemons/go-gypsy/yaml"
)

var (
	ErrNotYamlMap = errors.New("Cannot create task as passed node is not a YAML MAP")
)

type Task struct {
	Name         string
	Command      CommandIntf
	Tags         []string
	IgnoreErrors bool
	RegisterVar  string
	ChangedValue string
}

func NewTaskFromYaml(role *Role, node yaml.Node) (*Task, error) {
	var g yaml.Map
	var ok bool

	if g, ok = node.(yaml.Map); !ok {
		return nil, ErrNotYamlMap
	}

	task := &Task{}

	for k, v := range g {
		log.Printf("k = %s, v = %v", k, v)
		s := ""
		if q, ok := v.(yaml.Scalar); ok {
			s = string(q)
		}
		switch k {
		case "include":
			err := processRoleFile(role, string(s))
			if err != nil {
				return nil, err
			}
		case "ignore_errors":
			b, err := strconv.ParseBool(s)
			if err != nil {
				return nil, fmt.Errorf("Invalid 'boolean': %s", s)
			}
			task.IgnoreErrors = b
		case "name":
			task.Name = s
		case "tags":
			task.Tags = strings.Split(s, " ")
		case "notify":
			task.Tags = strings.Split(s, " ")
		case "when":
			// do something
		case "register":
			task.RegisterVar = s
		case "changed_when":
			task.ChangedValue = s
		case "args":
		case "with_items":
		default:
			if task.Command == nil {
				if err := CommandLookup(k); err != nil {
					return nil, fmt.Errorf("No such command: %s", k)
				}
				//							task.Command := core.CommandCreate(s)
				//				task.Command = true
			} else {
				return nil, fmt.Errorf("Cannot have multiple commands on each task: %s", k)
			}
		}
	}
	return task, nil
}
