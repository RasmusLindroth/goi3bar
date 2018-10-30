package memory

import (
	. "github.com/denbeigh2000/goi3bar"
	"github.com/denbeigh2000/goi3bar/config"

	"time"
)

const Identifier = "spotify"

// SpotifyConfig represents the configuration for a Memory plugin in
// JSON format
type SpotifyConfig struct {
	Interval string `json:"interval"`
}

type spotifyBuilder struct{}

func (m spotifyBuilder) Build(c config.Config) (Producer, error) {
	conf := SpotifyConfig{}
	err := c.ParseConfig(&conf)
	if err != nil {
		return nil, err
	}

	interval, err := time.ParseDuration(conf.Interval)
	if err != nil {
		return nil, err
	}

	g := Spotify{
		Name: Identifier,
	}

	return &BaseProducer{
		Generator: g,
		Name:      Identifier,
		Interval:  interval,
	}, nil
}

func init() {
	config.Register(Identifier, spotifyBuilder{})
}
