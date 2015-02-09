package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/Tomohiro/air/media"
	"github.com/codegangsta/cli"
	"github.com/gongo/go-airplay"
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
	path := c.Args().First()
	if path == "" {
		log.Fatal(fmt.Errorf("%s is not found", path))
	}

	mediaType, err := media.ClassifyType(path)
	if err != nil {
		log.Fatal(err)
	}

	var m media.Media

	switch mediaType {
	case media.IsDirectory:
		log.Fatal(errors.New("directory is not supported"))
	case media.IsFile:
		m = media.NewFile(path)
	}

	client, err := airplay.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	ch := client.Play(m.URL())
	<-ch
}
