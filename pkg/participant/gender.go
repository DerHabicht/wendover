package participant

import (
	"strings"

	"github.com/pkg/errors"
)

type Gender int

const (
	Male = iota
	Female
	NonBinary
)

func ParseGender(s string) (Gender, error) {
	switch strings.ToUpper(s) {
	case "MALE":
		return Male, nil
	case "FEMALE":
		return Female, nil
	case "NON-BINARY":
		return NonBinary, nil
	default:
		return -1, errors.Errorf("unrecognized gender: %s", s)
	}
}

func (g Gender) String() string {
	switch g {
	case Male:
		return "Male"
	case Female:
		return "Female"
	case NonBinary:
		return "Non-Binary"
	default:
		panic(errors.Errorf("invalid value for Gender: %d", g))
	}
}
