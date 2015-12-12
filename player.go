package main

import "github.com/gongo/go-airplay"

// Play plays the recieved paths of movie files to the Apple TV.
func Play(list []string) error {
	var err error

	// Initialize an AirPlay client device.
	client, err := airplay.FirstClient()
	if err != nil {
		return err
	}

	for _, path := range list {
		mediaType, err := classifyType(path)
		if err != nil {
			return err
		}

		var m media

		switch mediaType {
		case isFile:
			m = newFile(path)
		}

		ch := client.Play(m.URL())
		<-ch
	}

	return err
}
