package main

import "github.com/codegangsta/cli"

// Commands
var Commands = []cli.Command{
	commandDevices,
	commandPlay,
}

var commandDevices = cli.Command{
	Name:  "devices",
	Usage: "Show AirPlay devices",
	Description: `
`,
	Action: Devices,
}

var commandPlay = cli.Command{
	Name:  "play",
	Usage: "",
	Description: `
`,
	Action: Play,
}
