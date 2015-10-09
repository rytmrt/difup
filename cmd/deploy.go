package cmd

import (
	"fmt"
	flags "github.com/jessevdk/go-flags"
)

type Options struct {
	Dryrun bool   `short:"n" long:"dry-run" description:"Do dry-run"`
	Delete bool   `short:"d" long:"delete" description:"Do delete server file"`
	Branch string `short:"b" long:"branch" description:"diff branch" default:"master"`
}

type Deploy struct{}

func (f *Deploy) Help() string {
	return "Deploy"
}

func (f *Deploy) Run(args []string) int {
	var opts Options
	args, err := flags.Parse(&opts)

	if err != nil {
		return 1
	}

	if opts.Dryrun {
		fmt.Println("do dry-run!")
	}

	if opts.Delete {
		fmt.Println("server file delete")
	}

	fmt.Printf("%s\n", opts.Branch)

	return 0
}

func (f *Deploy) Synopsis() string {
	return "deploy"
}
