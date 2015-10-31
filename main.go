package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/gongo/go-airplay"
)

var (
	// exitCode to terminate.
	exitCode = 0
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	app := cli.NewApp()
	app.Name = "air"
	app.Version = Version
	app.Usage = "Command-line AirPlay client for Apple TV"
	app.Author = "Tomohiro TAIRA"
	app.Email = "tomohiro.t@gmail.com"
	app.Action = play
	app.Run(os.Args)
	return exitCode
}

func play(c *cli.Context) {
	path := c.Args().First()
	mediaType, err := classifyType(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		exitCode = 1
		return
	}

	var m media

	switch mediaType {
	case isFile:
		m = newFile(path)
	}

	client, err := airplay.FirstClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		exitCode = 1
		return
	}

	ch := client.Play(m.URL())
	<-ch
}
