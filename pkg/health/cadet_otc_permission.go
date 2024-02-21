package health

import (
	"github.com/google/uuid"
)

type CadetOTCPermission struct {
	id                uuid.UUID
	DisallowedOTCs    []OTC
	Reactions         string
	PermissionGranted bool
}
