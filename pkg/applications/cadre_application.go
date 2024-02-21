package applications

import (
	"github.com/google/uuid"
)

type CadreApplication struct {
	id             uuid.UUID
	Resume         Resume
	ClassesToTeach []string
	Preferences    []CadreRole
}
