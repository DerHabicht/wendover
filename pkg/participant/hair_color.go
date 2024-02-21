package participant

import (
	"strings"

	"github.com/pkg/errors"
)

type HairColor int

const (
	BaldHair = iota
	BlackHair
	BlondeHair
	BrownHair
	GrayHair
	RedHair
	SandyHair
	WhiteHair
)

func ParseHairColor(s string) (HairColor, error) {
	switch strings.ToUpper(s) {
	case "BALD":
		return BaldHair, nil
	case "BLACK":
		return BlackHair, nil
	case "BLONDE":
		return BlondeHair, nil
	case "BROWN":
		return BrownHair, nil
	case "GRAY":
		return GrayHair, nil
	case "RED":
		return RedHair, nil
	case "SANDY":
		return SandyHair, nil
	case "WHITE":
		return WhiteHair, nil
	default:
		return -1, errors.Errorf("unrecognized hair color: %s!", s)
	}
}

func (hc HairColor) String() string {
	switch hc {
	case BaldHair:
		return "Bald"
	case BlackHair:
		return "Black"
	case BlondeHair:
		return "Blonde"
	case BrownHair:
		return "Brown"
	case GrayHair:
		return "Gray"
	case RedHair:
		return "Red"
	case SandyHair:
		return "Sandy"
	case WhiteHair:
		return "White"
	default:
		panic(errors.Errorf("invalid value for HairColor: %d", hc))
	}
}
