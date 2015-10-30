package main

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/gongo/go-airplay"
)

func main() {
	newApp().Run(os.Args)
}

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
	mediaType, err := classifyType(path)
	if err != nil {
		log.Fatal(err)
	}

	var m media

	switch mediaType {
	case isFile:
		m = newFile(path)
	}

	client, err := airplay.FirstClient()
	if err != nil {
		log.Fatal(err)
	}

	ch := client.Play(m.URL())
	<-ch
}
