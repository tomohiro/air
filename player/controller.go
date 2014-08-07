package player

import (
	"errors"
	"fmt"

	"github.com/gongo/go-airplay"
)

// A Controller can play media files
type Controller struct {
	Playlist *Playlist
}

// NewController returns new controller
func NewController() *Controller {
	return new(Controller)
}

// SetPlaylist Set Playlist
func (c *Controller) SetPlaylist(p *Playlist) error {
	if len(p.Entries) == 0 {
		return errors.New("media files not found")
	}
	c.Playlist = p
	return nil
}

// Play all entries
func (c *Controller) Play() error {
	client, err := airplay.NewClient()
	if err != nil {
		return err
	}

	for _, media := range c.Playlist.Entries {
		fmt.Println(media.Path)
		ch := client.Play(source(media.Path))
		<-ch
	}
	return nil
}
