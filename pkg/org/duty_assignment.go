package org

import (
	"github.com/google/uuid"

	"github.com/derhabicht/wendover/pkg/participant"
)

type DutyAssignment struct {
	id                 uuid.UUID
	role               CadreRole
	title              string
	Assignee           *participant.Participant
	directSubordinates []DutyAssignment
}

func (da DutyAssignment) ID() uuid.UUID {
	return da.id
}

func (da DutyAssignment) Role() CadreRole {
	return da.role
}

func (da DutyAssignment) Title() string {
	return da.title
}

func (da DutyAssignment) DirectSubordinates() []DutyAssignment {
	return da.directSubordinates
}
