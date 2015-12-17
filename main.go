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
	app := cli.NewApp()
	app.Name = "air"
	app.Version = Version
	app.Usage = "Command-line AirPlay client for Apple TV"
	app.Author = "Tomohiro TAIRA"
	app.Email = "tomohiro.t@gmail.com"
	app.Action = func(c *cli.Context) {
		paths := c.Args()
		if !paths.Present() {
			fmt.Fprintf(os.Stderr, "Incorrect usage.\nRun `air <path>`\n")
			exitCode = 1
			return
		}

		if err := Play(paths); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			exitCode = 1
			return
		}
	}
	app.Run(os.Args)

	os.Exit(exitCode)
}

// Play plays the recieved paths of media files to the Apple TV.
func Play(paths []string) error {
	var err error

	// Initialize an AirPlay client device.
	client, err := airplay.FirstClient()
	if err != nil {
		return err
	}

	for _, path := range paths {
		url, err := Open(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Skipped %s. because %s.\n", path, err)
			continue
		}

		ch := client.Play(url)
		<-ch
	}

	return err
}
