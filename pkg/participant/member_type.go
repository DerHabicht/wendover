package participant

import (
	"strings"

	"github.com/pkg/errors"
)

type MemberType int

const (
	SeniorMember = iota
	CadetSponsorMember
	CadetMember
)

func ParseMemberType(s string) (int, error) {
	switch strings.ToUpper(s) {
	case "SENIOR":
		return SeniorMember, nil
	case "CADET SPONSOR":
		return CadetSponsorMember, nil
	case "CADET":
		return CadetMember, nil
	default:
		return -1, errors.Errorf("invalid member type: %s", s)
	}
}

func (mt MemberType) String() string {
	switch mt {
	case SeniorMember:
		return "SENIOR"
	case CadetSponsorMember:
		return "CADET SPONSOR"
	case CadetMember:
		return "CADET"
	default:
		panic(errors.Errorf("invalid value for MemberType: %d", mt))
	}
}
