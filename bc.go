package main

import (
	// Command plugins

	"io/ioutil"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/patpatnz/BlackCrystal/internal/core"
	"github.com/patpatnz/BlackCrystal/internal/hosts"

	_ "github.com/patpatnz/BlackCrystal/internal/cmds/action"
	_ "github.com/patpatnz/BlackCrystal/internal/cmds/apt"
	_ "github.com/patpatnz/BlackCrystal/internal/cmds/assert"
	_ "github.com/patpatnz/BlackCrystal/internal/cmds/copy"
	_ "github.com/patpatnz/BlackCrystal/internal/cmds/docker"
	_ "github.com/patpatnz/BlackCrystal/internal/cmds/file"
	_ "github.com/patpatnz/BlackCrystal/internal/cmds/get_url"
	_ "github.com/patpatnz/BlackCrystal/internal/cmds/hostname"
	_ "github.com/patpatnz/BlackCrystal/internal/cmds/lineinfile"
	_ "github.com/patpatnz/BlackCrystal/internal/cmds/pip"
	_ "github.com/patpatnz/BlackCrystal/internal/cmds/service"
	_ "github.com/patpatnz/BlackCrystal/internal/cmds/shell"
	_ "github.com/patpatnz/BlackCrystal/internal/cmds/template"
	_ "github.com/patpatnz/BlackCrystal/internal/cmds/user"
)

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

func main() {

	myhosts, err := hosts.NewFromFile("/Users/pjs/GO/src/github.com/bnsl/buddyguard/ansible/hosts")
	if err != nil {
		log.Fatal(err)
	}

	_ = myhosts

	/*	err = loadRoles()
		if err != nil {
			log.Print(err)
			return
		}*/
}
