package participant

import (
	"strings"

	"github.com/pkg/errors"
)

type TShirtSize int

const (
	ShirtXXXL = iota
	ShirtXXL
	ShirtXL
	ShirtL
	ShirtM
	ShirtS
	ShirtXS
	ShirtYL
	ShirtYM
)

func ParseTShirtSize(s string) (TShirtSize, error) {
	switch strings.ToUpper(s) {
	case "XXXL":
		return ShirtXXXL, nil
	case "XXL":
		return ShirtXXL, nil
	case "XL":
		return ShirtXL, nil
	case "L":
		return ShirtL, nil
	case "M":
		return ShirtM, nil
	case "S":
		return ShirtS, nil
	case "XS":
		return ShirtXS, nil
	case "YL":
		return ShirtYL, nil
	case "YM":
		return ShirtYM, nil
	default:
		return -1, errors.Errorf("invalid T-shirt size: %s", s)
	}
}

func (tss TShirtSize) String() string {
	switch tss {
	case ShirtXXXL:
		return "XXXL"
	case ShirtXXL:
		return "XXL"
	case ShirtXL:
		return "XL"
	case ShirtL:
		return "L"
	case ShirtM:
		return "M"
	case ShirtS:
		return "S"
	case ShirtXS:
		return "XS"
	case ShirtYL:
		return "YL"
	case ShirtYM:
		return "YM"
	default:
		panic(errors.Errorf("invalid value for TShirtSize: %d", tss))
	}
}
