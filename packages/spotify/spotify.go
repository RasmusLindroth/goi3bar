package memory

import (
	i3 "github.com/denbeigh2000/goi3bar"
	"github.com/godbus/dbus"

	"fmt"
)

var bus *dbus.BusObject

func newBus() *dbus.BusObject {

	conn, err := dbus.SessionBus()

	// couldnt connect to session bus
	if err != nil {
		panic(err)
	}

	obj := conn.Object("org.mpris.MediaPlayer2.spotify", "/org/mpris/MediaPlayer2")
	return &obj
}

type Song struct {
	Artist string
	Title  string
}

func (s *Song) Update() {
	if bus == nil {
		bus = newBus()
	}
	b := *bus
	meta, err := b.GetProperty("org.mpris.MediaPlayer2.Player.Metadata")

	if err != nil {
		s.Artist = ""
		s.Title = ""
		return
	}

	songData, ok := meta.Value().(map[string]dbus.Variant)
	if !ok {
		return
	}
	if artist, ok := songData["xesam:artist"]; ok {
		tmp := artist.Value().([]string)
		if len(tmp) < 1 {
			return
		}
		s.Artist = artist.Value().([]string)[0]
	}
	if title, ok := songData["xesam:title"]; ok {
		s.Title = title.Value().(string)
	}

}

func (s *Song) Get() string {
	return fmt.Sprintf("%v - %v", s.Artist, s.Title)
}

type Spotify struct {
	Name string
}

func (s Spotify) Generate() ([]i3.Output, error) {

	song := Song{}
	song.Update()
	current := song.Get()
	output := current

	var color string
	switch current {
	case " - ":
		color = i3.DefaultColors.Crit
		output = "No music"
	default:
		color = i3.DefaultColors.OK
	}
	out := make([]i3.Output, 1)

	out[0] = i3.Output{
		Name:      s.Name,
		FullText:  output,
		Color:     color,
		Separator: true,
	}

	return out, nil
}
