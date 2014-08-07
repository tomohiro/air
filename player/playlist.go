package player

import (
	"errors"
	"fmt"

	"github.com/codegangsta/cli"

	"os"
	"path/filepath"
)

// Playlist have multiple media file
type Playlist struct {
	Entries []*Media
}

// NewPlaylist creates a new playlist
func NewPlaylist() *Playlist {
	return new(Playlist)
}

// Add media file to playlist
func (p *Playlist) Add(c *cli.Context) error {
	path := c.Args().First()
	if path == "" {
		return fmt.Errorf("%s is not found", path)
	}

	path, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return err
	}

	stat, err := f.Stat()
	if err != nil {
		return err
	}

	switch mode := stat.Mode(); {
	case mode.IsDir():
		fmt.Println(f.Name() + " is directory")
		return errors.New("directory is not supported")
	case mode.IsRegular():
		p.Entries = append(p.Entries, NewMedia(path))
	}
	return nil
}
