package main

import (
	// Command plugins
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/kylelemons/go-gypsy/yaml"
	"github.com/patpatnz/BlackCrystal/internal/core"

	_ "github.com/patpatnz/BlackCrystal/internal/cmds/assert"
	_ "github.com/patpatnz/BlackCrystal/internal/cmds/copy"
	_ "github.com/patpatnz/BlackCrystal/internal/cmds/file"
	_ "github.com/patpatnz/BlackCrystal/internal/cmds/template"
	_ "github.com/patpatnz/BlackCrystal/internal/cmds/user"
)

type Task struct {
	Name    string
	Command bool //core.CommandIntf
	Tags    []string
}

type Role struct {
	Tasks []*Task
}

func processRoleFile(role *Role, dir, fileName string) error {

	log.Print(fileName)
	dat, err := os.Open(dir + "/" + fileName)
	if err != nil {
		return err
	}

	v, err := yaml.Parse(dat)
	if w, ok := v.(yaml.List); ok {
		for _, y := range w {
			if g, ok := y.(yaml.Map); ok {
				task := &Task{}
				add := true
				for k, v := range g {
					log.Printf("k = %s, v = %v", k, v)
					s := ""
					if q, ok := v.(yaml.Scalar); ok {
						s = string(q)
					}
					switch k {
					case "include":
						err = processRoleFile(role, dir, string(s))
						if err != nil {
							return err
						}
						add = false
					case "name":
						task.Name = s
					case "tags":
						task.Tags = strings.Split(s, " ")
					case "notify":
						task.Tags = strings.Split(s, " ")
					default:
						if !task.Command {
							if err := core.CommandLookup(k); err != nil {
								return fmt.Errorf("No such command: %s", k)
							}
							//							task.Command := core.CommandCreate(s)
							task.Command = true
						} else {
							return fmt.Errorf("Blarg: %s", k)
						}
					}
				}
				if add {
					role.Tasks = append(role.Tasks, task)
				}
			}
		}
	}
	log.Printf("Finished %s", fileName)

	return nil
}

func loadRoles() error {
	baseDIR := "/Users/pjs/Projects/GO/src/github.com/bnsl/buddyguard/ansible/roles"

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

		role := &Role{Tasks: make([]*Task, 0)}
		err := processRoleFile(role, startFile, "main.yml")
		spew.Dump(role)
		if err != nil {
			return err
		}

		//for _, y := range v {
		//	_ = y
		//}

		//		spew.Dump(v)

		return nil
	}

	return nil
}

func main() {
	err := core.CommandLookup("user")
	if err != nil {
		log.Print(err)
	}
	/*
		b := make(map[string]interface{})
		in, _ := ioutil.ReadFile("script.yaml")

		yaml.Unmarshal(in, b)
	*/
	err = loadRoles()
	if err != nil {
		log.Print(err)
		return
	}
}
