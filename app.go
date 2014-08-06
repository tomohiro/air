package main

import (
	"fmt"
	"log"

	"github.com/codegangsta/cli"
	"github.com/gongo/go-airplay"
)

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "air"
	app.Version = Version
	app.Usage = "Command-line AirPlay video client for Apple TV"
	app.Author = "Tomohiro TAIRA"
	app.Email = "tomohiro.t@gmail.com"
	app.Commands = []cli.Command{
		{
			Name:   "play",
			Usage:  "Play media file(Movie, Music)",
			Action: play,
		},
		{
			Name:   "devices",
			Usage:  "Show AirPlay devices",
			Action: devices,
		},
	}
	return app
}

func play(c *cli.Context) {
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
}

func devices(c *cli.Context) {
	devices, err := Devices()
	if err != nil {
		log.Fatal(err)
	}
	for _, d := range devices {
		fmt.Println(d.Name)
	}
}
