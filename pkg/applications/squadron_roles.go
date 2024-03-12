package applications

import (
	"github.com/pkg/errors"
)

type SquadronRole int

const (
	SqdnCadetCommander = iota
	SqdnCadetDeputyCommander
	SqdnFirstSergeant
	SqdnCadetFlightCommander
	SqdnFlightSergeant
	SqdnElementLeader
	SqdnSupportStaff
)

func ParseSquadronRole(s string) SquadronRole {
	panic(errors.New("method not implemented"))
}

func (sr SquadronRole) String() string {
	panic(errors.New("method not implemented"))
}

func (sr SquadronRole) Scan() {
	panic(errors.New("method not implemented"))
}

func (sr SquadronRole) Value() {
	panic(errors.New("method not implemented"))
}
