package main

import (
  "fmt"

  "github.com/codegangsta/cli"
  "github.com/gongo/go-airplay"
)

// Show devices
func Devices(c *cli.Context) {
  for _, d := range airplay.Devices() {
    fmt.Println(d.Name)
  }
}
