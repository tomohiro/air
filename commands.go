package main

import (
	"github.com/codegangsta/cli"
	"github.com/gongo/go-airplay"
)

// Commands
var Commands = []cli.Command{
	commandPlay,
}

var commandPlay = cli.Command{
	Name:  "play",
	Usage: "",
	Description: `
`,
	Action: doPlay,
}

func doPlay(c *cli.Context) {
	client := airplay.NewClient()
	ch := client.Play(c.Args().First())
	<-ch
}
