package applications

import (
	"github.com/google/uuid"

	"github.com/derhabicht/wendover/pkg/flights"
	"github.com/derhabicht/wendover/pkg/health"
	"github.com/derhabicht/wendover/pkg/participant"
)

type Application struct {
	id                 uuid.UUID
	participant        participant.Participant
	contactInformation participant.ContactInformation
	flightInformation  *flights.FlightInformation
	cadreApplication   CadreApplication
	healthInformation  health.Information
	emergencyContact   participant.ContactInformation
}

func NewApplication(
	id uuid.UUID,
	participant participant.Participant,
	contactInformation participant.ContactInformation,
	flightInformation *flights.FlightInformation,
	cadreApplication CadreApplication,
	healthInformation health.Information,
	emergencyContact participant.ContactInformation,
) Application {
	return Application{
		id:                 id,
		participant:        participant,
		contactInformation: contactInformation,
		flightInformation:  flightInformation,
		cadreApplication:   cadreApplication,
		healthInformation:  healthInformation,
		emergencyContact:   emergencyContact,
	}
}

func (a *Application) ID() uuid.UUID {
	return a.id
}
