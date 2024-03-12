package applications

import (
	"github.com/google/uuid"

	"github.com/derhabicht/wendover/pkg/org"
)

type Resume struct {
	id                  uuid.UUID
	encampmentsAttended uint
	squadronRoleHistory []SquadronRole
	cadreRoleHistory    []org.CadreRole
}
