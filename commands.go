package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/gongo/go-airplay"
)

// Define commands
var Commands = []cli.Command{
	{
		Name:  "play",
		Usage: "Play media file(Movie, Music)",
		Action: func(c *cli.Context) {
			client, err := airplay.NewClient()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
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
