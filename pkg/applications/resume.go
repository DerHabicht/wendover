package applications

import (
	"github.com/google/uuid"
)

type Resume struct {
	id                  uuid.UUID
	EncampmentsAttended uint
	SquadronRoleHistory []SquadronRole
	CadreRoleHistory    []CadreRole
}
