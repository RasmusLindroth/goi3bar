package memory

import (
	"os/exec"
	"strings"

	i3 "github.com/denbeigh2000/goi3bar"

	"fmt"
)

var servers = map[string]string{
	"mullvad-au1":  "Australia",
	"mullvad-at1":  "Austria",
	"mullvad-be1":  "Belgium",
	"mullvad-br1":  "Brazil",
	"mullvad-bg1":  "Bulgaria",
	"mullvad-ca3":  "Canada",
	"mullvad-ca1":  "Canada",
	"mullvad-ca2":  "Canada",
	"mullvad-cz1":  "Czech Republic",
	"mullvad-dk1":  "Denmark",
	"mullvad-fi1":  "Finland",
	"mullvad-fr1":  "France",
	"mullvad-de1":  "Germany",
	"mullvad-de2":  "Germany",
	"mullvad-de4":  "Germany",
	"mullvad-de5":  "Germany",
	"mullvad-hk1":  "Hong Kong",
	"mullvad-in1":  "India",
	"mullvad-it1":  "Italy",
	"mullvad-jp1":  "Japan",
	"mullvad-md1":  "Moldova",
	"mullvad-nl1":  "Netherlands",
	"mullvad-nl2":  "Netherlands",
	"mullvad-nl3":  "Netherlands",
	"mullvad-no1":  "Norway",
	"mullvad-pl1":  "Poland",
	"mullvad-ro1":  "Romania",
	"mullvad-rs1":  "Serbia",
	"mullvad-sg1":  "Singapore",
	"mullvad-es1":  "Spain",
	"mullvad-se3":  "Gothenburg",
	"mullvad-se5":  "Gothenburg",
	"mullvad-se4":  "Malmö",
	"mullvad-se2":  "Stockholm",
	"mullvad-se6":  "Stockholm",
	"mullvad-se7":  "Stockholm",
	"mullvad-se8":  "Stockholm",
	"mullvad-ch1":  "Switzerland",
	"mullvad-ch2":  "Switzerland",
	"mullvad-gb2":  "UK",
	"mullvad-gb4":  "UK",
	"mullvad-gb5":  "UK",
	"mullvad-gb3":  "UK",
	"mullvad-us6":  "USA",
	"mullvad-us4":  "USA",
	"mullvad-us7":  "USA",
	"mullvad-us11": "USA",
	"mullvad-us12": "USA",
	"mullvad-us2":  "USA",
	"mullvad-us3":  "USA",
	"mullvad-us1":  "USA",
	"mullvad-us13": "USA",
	"mullvad-us15": "USA",
	"mullvad-us14": "USA",
	"mullvad-us9":  "USA",
	"mullvad-us5":  "USA",
	"mullvad-ua1":  "Ukraine",
}

func getMullvad() string {
	res, err := exec.Command("ip", "link").Output()

	if err != nil {
		return "VPN: ?"
	}

	lines := strings.Split(string(res), "\n")

	var vpns []string
	for _, line := range lines {
		parts := strings.Split(line, " ")

		if len(parts) < 2 || len(parts[1]) < 11 {
			continue
		}

		intface := parts[1][0 : len(parts[1])-1]
		if intface[:7] != "mullvad" {
			continue
		}

		if val, ok := servers[intface]; ok {
			key := strings.Split(intface, "-")[1]
			intface = fmt.Sprintf("%s (%s)", val, key)
		}

		vpns = append(vpns, intface)
	}

	vpn := strings.Join(vpns, " - ")
	if vpn == "" {
		vpn = "down"
	}
	return fmt.Sprintf("VPN: %s", vpn)
}

type Mullvad struct {
	Name string
}

const (
	FormatString = "%v"
)

func (m Mullvad) Generate() ([]i3.Output, error) {

	mullvadStatus := getMullvad()

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
