package health

import (
	"github.com/derhabicht/wendover/pkg/participant"
)

type Information struct {
	History              History
	MedicalInsurance     Insurance
	PrescriptionCoverage Insurance
	PhysicianContactInfo participant.ContactInformation
}
