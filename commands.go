package main

import (
	"fmt"
	"log"

	"github.com/codegangsta/cli"
	"github.com/gongo/go-airplay"
)

// Define commands
var Commands = []cli.Command{
	{
		Name:  "play",
		Usage: "Play media file(Movie, Music)",
		Action: func(c *cli.Context) {
			target := c.Args().First()
			playlist := NewPlaylist()
			err := playlist.Add(target)
			if err != nil {
				log.Fatal(err)
			}

			client, err := airplay.NewClient()
			if err != nil {
				log.Fatal(err)
			}
			for _, media := range playlist.Entries {
				fmt.Println(media.Path)
				ch := client.Play(source(media.Path))
				<-ch
			}
		},
	},
	{
		Name:  "devices",
		Usage: "Show AirPlay devices",
		Action: func(c *cli.Context) {
			devices, err := Devices()
			if err != nil {
				log.Fatal(err)
			}
			for _, d := range devices {
				fmt.Println(d.Name)
			}
		},
	},
}
