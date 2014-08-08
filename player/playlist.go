package player

import (
	"errors"
	"fmt"

	"github.com/Tomohiro/air/media"
	"github.com/codegangsta/cli"
)

// Playlist have multiple media file
type Playlist struct {
	Entries []media.Media
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

	mediaType, err := media.ClassifyType(path)
	if err != nil {
		return err
	}

	switch mediaType {
	case media.IsDirectory:
		fmt.Println(path + " is directory")
		return errors.New("directory is not supported")
	case media.IsFile:
		p.Entries = append(p.Entries, media.NewFile(path))
	}
	return nil
}
