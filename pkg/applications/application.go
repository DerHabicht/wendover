package applications

import (
	"github.com/google/uuid"

	"github.com/derhabicht/wendover/pkg/flights"
	"github.com/derhabicht/wendover/pkg/health"
	"github.com/derhabicht/wendover/pkg/participant"
)

type Application struct {
	id                 uuid.UUID
	Participant        participant.Participant
	ContactInformation participant.ContactInformation
	FlightInformation  *flights.FlightInformation
	CadreApplication   CadreApplication
	HealthInformation  health.Information
	EmergencyContact   participant.ContactInformation
}
