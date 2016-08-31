package core

import (
	"log"
	"os"
	"path/filepath"

	"github.com/kylelemons/go-gypsy/yaml"
)

type Role struct {
	Name  string
	Tasks []*Task
	Dir   string
}

func NewRoleFromFile(fileName string) (*Role, error) {
	dir, _ := filepath.Abs(filepath.Dir(fileName) + "/..")
	file := filepath.Base(fileName)

	role := &Role{Tasks: make([]*Task, 0), Dir: dir}

	err := processRoleFile(role, file)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func processRoleFile(role *Role, fileName string) error {
	log.Print(fileName)
	dat, err := os.Open(role.Dir + "/tasks/" + fileName)
	if err != nil {
		return err
	}

	v, err := yaml.Parse(dat)
	if w, ok := v.(yaml.List); ok {
		for _, y := range w {
			add := true
			task, err := NewTaskFromYaml(role, y)
			if err != nil {
				return err
			}
			if add {
				role.Tasks = append(role.Tasks, task)
			}
		}
	}
	log.Printf("Finished %s", fileName)

	return nil
}
