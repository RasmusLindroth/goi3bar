package memory

import (
	. "github.com/denbeigh2000/goi3bar"
	"github.com/denbeigh2000/goi3bar/config"

	"time"
)

const Identifier = "mullvad"

// MullvadConfig represents the configuration for a Memory plugin in
// JSON format
type MullvadConfig struct {
	Interval string `json:"interval"`
}

type mullvadBuilder struct{}

func (m mullvadBuilder) Build(c config.Config) (Producer, error) {
	conf := MullvadConfig{}
	err := c.ParseConfig(&conf)
	if err != nil {
		return nil, err
	}

	interval, err := time.ParseDuration(conf.Interval)
	if err != nil {
		return nil, err
	}

	g := Mullvad{
		Name: Identifier,
	}

	return &BaseProducer{
		Generator: g,
		Name:      Identifier,
		Interval:  interval,
	}, nil
}

func init() {
	config.Register(Identifier, mullvadBuilder{})
}
