package applications

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
