package cmd

import ()

type Deploy struct{}

func (f *Deploy) Help() string {
	return "Deploy"
}

func (f *Deploy) Run(args []string) int {
	return 0
}

func (f *Deploy) Synopsis() string {
	return "deploy"
}
