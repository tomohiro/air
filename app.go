package main

import (
	"fmt"
	"log"

	"github.com/Tomohiro/air/player"
	"github.com/codegangsta/cli"
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
			Name:  "play",
			Usage: "Play media file(Movie, Music)",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "dir, d",
					Value: "",
					Usage: "directory",
				},
				cli.StringFlag{
					Name:  "file, f",
					Value: "",
					Usage: "file",
				},
			},
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
	playlist := player.NewPlaylist()
	if err := playlist.Add(c); err != nil {
		log.Fatal(err)
	}

	controller := player.NewController()
	if err := controller.SetPlaylist(playlist); err != nil {
		log.Fatal(err)
	}

	if err := controller.Play(); err != nil {
		log.Fatal(err)
	}
}

func devices(c *cli.Context) {
	devices, err := player.Devices()
	if err != nil {
		log.Fatal(err)
	}
	for _, d := range devices {
		fmt.Println(d.Name)
	}
}
