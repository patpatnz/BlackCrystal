package main

import (
	// Command plugins

	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/patpatnz/BlackCrystal/internal/hosts"
	"github.com/patpatnz/BlackCrystal/internal/hostvars"
	"github.com/patpatnz/BlackCrystal/internal/job"
	"github.com/patpatnz/BlackCrystal/internal/transport"

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

func main() {

	myhosts, err := hosts.NewFromFile("hosts")
	if err != nil {
		log.Fatal(err)
	}

	h := myhosts.GetHosts("patdb")
	spew.Dump(h)

	tp, err := transport.Get("ssh", h[0])
	if err != nil {
		return
	}

	j := &job.Job{Transport: tp, Host: h[0]}

	hostvars.Run(j)

	/*	err = loadRoles()
		if err != nil {
			log.Print(err)
			return
		}*/
}
