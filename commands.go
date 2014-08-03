package main

import "github.com/codegangsta/cli"

// Commands
var Commands = []cli.Command{
	commandPlay,
}

var commandPlay = cli.Command{
	Name:  "play",
	Usage: "",
	Description: `
`,
	Action: Play,
}
