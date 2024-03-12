package applications

import (
	"github.com/google/uuid"

	"github.com/derhabicht/wendover/pkg/org"
)

type CadreApplication struct {
	id             uuid.UUID
	resume         Resume
	classesToTeach []string
	preferences    []org.CadreRole
}
