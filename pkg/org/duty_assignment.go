package org

import (
	"github.com/google/uuid"

	"github.com/derhabicht/wendover/pkg/applications"
	"github.com/derhabicht/wendover/pkg/participant"
)

type DutyAssignment struct {
	id           uuid.UUID
	role         applications.CadreRole
	dutyTitle    string
	assignee     *participant.Participant
	subordinates []DutyAssignment
}
