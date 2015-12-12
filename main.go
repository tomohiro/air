package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
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
	path := c.Args()
	if len(path) == 0 {
		fmt.Fprintf(os.Stderr, "Incorrect usage.\nRun `air <path>`\n")
		exitCode = 1
		return
	}

	if err := Play(path); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		exitCode = 1
		return
	}
}
