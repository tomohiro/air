package main

import (
	"fmt"
	"os"

	"github.com/gongo/go-airplay"
)

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
