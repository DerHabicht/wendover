package participant

import (
	"net/mail"

	"github.com/google/uuid"
	"github.com/nyaruka/phonenumbers"
)

type ContactInformation struct {
	id                   uuid.UUID
	contactName          string
	emailAddress         mail.Address
	primaryPhoneNumber   phonenumbers.PhoneNumber
	secondaryPhoneNumber *phonenumbers.PhoneNumber
	mailingAddress       *string
}
