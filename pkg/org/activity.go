package org

import (
	"time"

	"github.com/google/uuid"
)

type Activity struct {
	id                uuid.UUID
	Key               string
	Name              string
	Location          string
	Start             time.Time
	End               time.Time
	CadetStudentFee   uint
	CadetCadreFee     uint
	SeniorStudentFee  uint
	SeniorCadreFee    uint
	ActivityCommander DutyAssignment
	ActivityUnits     []ActivityUnit
}

func (a Activity) ID() uuid.UUID {
	return a.id
}
