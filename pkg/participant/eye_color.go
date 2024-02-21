package participant

import (
	"strings"

	"github.com/pkg/errors"
)

type EyeColor int

const (
	BlackEyes = iota
	BlueEyes
	BrownEyes
	GrayEyes
	GreenEyes
	HazelEyes
	MaroonEyes
)

func ParseEyeColor(s string) (EyeColor, error) {
	switch strings.ToUpper(s) {
	case "BLACK":
		return BlackEyes, nil
	case "BLUE":
		return BlueEyes, nil
	case "BROWN":
		return BrownEyes, nil
	case "GRAY":
		return GrayEyes, nil
	case "GREEN":
		return GreenEyes, nil
	case "HAZEL":
		return HazelEyes, nil
	case "MAROON":
		return MaroonEyes, nil
	default:
		return -1, errors.Errorf("unrecognized eye color: %s", s)
	}
}

func (ec EyeColor) String() string {
	switch ec {
	case BlackEyes:
		return "Black"
	case BlueEyes:
		return "Blue"
	case BrownEyes:
		return "Brown"
	case GrayEyes:
		return "Gray"
	case GreenEyes:
		return "Green"
	case HazelEyes:
		return "Hazel"
	case MaroonEyes:
		return "Maroon"
	default:
		panic(errors.Errorf("invalid value for EyeColor: %d", ec))
	}
}
