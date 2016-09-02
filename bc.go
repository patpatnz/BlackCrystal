package main

import (
	// Command plugins

	"io/ioutil"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/patpatnz/BlackCrystal/internal/core"
	"github.com/patpatnz/BlackCrystal/internal/hosts"
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

	myhosts, err := hosts.NewFromFile("/Users/pjs/Projects/GO/src/github.com/bnsl/buddyguard/ansible/hosts")
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
