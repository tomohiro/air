package main

import (
	"log"

	"github.com/Tomohiro/air/player"
	"github.com/codegangsta/cli"
)

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "air"
	app.Version = Version
	app.Usage = "Command-line AirPlay client for Apple TV"
	app.Author = "Tomohiro TAIRA"
	app.Email = "tomohiro.t@gmail.com"
	app.Action = play
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
