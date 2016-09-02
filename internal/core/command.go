package core

import (
	"errors"
	"strings"

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

// CommandIntf defines the interface that command types must implement
type CommandIntf interface {
	Setup(params []CommandParameter) error
}

type CommandInstance func() (CommandIntf, error)

// Command structure defines a executable command
type Command struct {
	Name        string
	Parameters  []CommandParameter
	GetInstance CommandInstance
}

// CommandParameter fdksjfdks
type CommandParameter struct {
	Name  string
	Type  CommandParameterType
	Value interface{}
}

// CommandParameterType represents the type of value
type CommandParameterType int

const (
	// CommandParameterInt is an integer value
	CommandParameterInt CommandParameterType = iota
	// CommandParameterString is a string value
	CommandParameterString
)

var (
	// ErrCommandNotFound is an error returned when the command doesnt exist
	ErrCommandNotFound = errors.New("No such command.")
)

var (
	commands = make(map[string]*Command)
)

func CommandRegister(cmd Command) {
	if _, ok := commands[strings.ToLower(cmd.Name)]; ok {
		panic("A double registration error occured!")
	}
	commands[strings.ToLower(cmd.Name)] = &cmd
}

func CommandLookup(name string) error {
	if _, ok := commands[strings.ToLower(name)]; ok {
		return nil
	}
	return ErrCommandNotFound
}
