package main

import (
	"github.com/mitchellh/cli"
	"github.com/rytmrt/difup/cmd"
	"log"
)

func Run(args []string) int {
	c := cli.NewCLI(COMMAND, VERSION)

	c.Args = args[0:]

	c.Commands = map[string]cli.CommandFactory{
		"deploy": func() (cli.Command, error) {
			return &cmd.Deploy{}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	return exitStatus
}
