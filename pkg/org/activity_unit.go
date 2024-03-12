package org

import (
	"github.com/google/uuid"

	"github.com/derhabicht/wendover/pkg/participant"
)

type ActivityUnit struct {
	id               uuid.UUID
	Name             string
	Commander        DutyAssignment
	SubordinateUnits []ActivityUnit
	Participants     []participant.Participant
}

func (au ActivityUnit) ID() uuid.UUID {
	return au.id
}
