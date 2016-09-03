package core

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/davecgh/go-spew/spew"
	"github.com/kylelemons/go-gypsy/yaml"
	"github.com/patpatnz/BlackCrystal/internal/core"
)

type Role struct {
	Name  string
	Tasks []*Task
	Dir   string
}

func loadRoles() error {
	baseDIR := "/Users/pjs/GO/src/github.com/bnsl/buddyguard/ansible/roles"

	rolefiles, err := ioutil.ReadDir(baseDIR)
	if err != nil {
		return err
	}
	for _, v := range rolefiles {
		if v.Name() == "." || v.Name() == ".." {
			continue
		}
		log.Printf("Loading role %s:", v.Name())
		startFile := baseDIR + "/" + v.Name() + "/tasks"

		role, err := core.NewRoleFromFile(startFile + "/main.yml")
		spew.Dump(role)
		if err != nil {
			return err
		}
	}

	return nil
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
