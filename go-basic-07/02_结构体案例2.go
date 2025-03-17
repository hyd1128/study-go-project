package main

import "fmt"

type Command struct {
	Name    string
	Var     *int
	Comment string
}

func newCommand(name string, varRef *int, command string) *Command {
	return &Command{
		Name:    name,
		Var:     varRef,
		Comment: command,
	}
}

var version = 1

func main() {
	cmd := newCommand(
		"version",
		&version,
		"show version",
	)
	fmt.Println(cmd)
}
