package org

import (
	"github.com/google/uuid"

	"github.com/derhabicht/wendover/pkg/participant"
)

type CadreRole struct {
	id                uuid.UUID
	title             string
	memberType        participant.MemberType
	minGrade          *participant.Grade
	maxGrade          *participant.Grade
	preferredMinGrade *participant.Grade
	preferredMaxGrade *participant.Grade
}

func (cr CadreRole) ID() uuid.UUID {
	return cr.id
}

func (cr CadreRole) Title() string {
	return cr.title
}

func (cr CadreRole) MemberType() participant.MemberType {
	return cr.memberType
}

func (cr CadreRole) MinGrade() *participant.Grade {
	return cr.minGrade
}

func (cr CadreRole) MaxGrade() *participant.Grade {
	return cr.maxGrade
}

func (cr CadreRole) PreferredMinGrade() *participant.Grade {
	return cr.preferredMinGrade
}

func (cr CadreRole) PreferredMaxGrade() *participant.Grade {
	return cr.preferredMaxGrade
}
