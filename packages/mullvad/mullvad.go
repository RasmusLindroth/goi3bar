package memory

import (
	"encoding/json"
	"net/http"
	"time"

	i3 "github.com/denbeigh2000/goi3bar"

	"fmt"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJSON(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func callMullvad() string {
	m := MullvadAPI{}
	err := getJSON("https://am.i.mullvad.net/json", &m)

	if err != nil {
		return "VPN ?"
	}

	if m.MullvadExitIP == true {
		return m.MullvadHostname
	}
	return "VPN down"

}

type MullvadAPI struct {
	IP                string  `json:"ip"`
	Country           string  `json:"country"`
	City              string  `json:"city"`
	Longitude         float64 `json:"longitude"`
	Latitude          float64 `json:"latitude"`
	MullvadExitIP     bool    `json:"mullvad_exit_ip"`
	MullvadHostname   string  `json:"mullvad_exit_ip_hostname"`
	Organization      string  `json:"organization"`
	MullvadServerType string  `json:"mullvad_server_type"`
	Blacklisted       struct {
		Blacklisted bool `json:"blacklisted"`
		Results     []struct {
			Name        string `json:"name"`
			Link        string `json:"link"`
			Blacklisted bool   `json:"blacklisted"`
		} `json:"results"`
	} `json:"blacklisted"`
}

type Mullvad struct {
	Name string
}

const (
	FormatString = "%v"
)

func (m Mullvad) Generate() ([]i3.Output, error) {

	mullvadStatus := callMullvad()

	var color string
	switch mullvadStatus {
	case "VPN ?":
		color = i3.DefaultColors.Warn
	case "VPN down":
		color = i3.DefaultColors.Crit
	default:
		color = i3.DefaultColors.OK
	}
	out := make([]i3.Output, 1)

	out[0] = i3.Output{
		Name:      m.Name,
		FullText:  fmt.Sprintf(FormatString, mullvadStatus),
		Color:     color,
		Separator: true,
	}

	return out, nil
}
