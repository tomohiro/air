package main

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/gongo/go-airplay"
)

// Define commands
var Commands = []cli.Command{
	{
		Name:  "play",
		Usage: "Play media file",
		Action: func(c *cli.Context) {
			client := airplay.NewClient()
			ch := client.Play(source(c.Args().First()))
			<-ch
		},
	},
	{
		Name:  "devices",
		Usage: "Show AirPlay devices",
		Action: func(c *cli.Context) {
			for _, d := range Devices() {
				fmt.Println(d.Name)
			}
		},
	},
}
