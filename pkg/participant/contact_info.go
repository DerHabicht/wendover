package participant

import (
	"net/mail"

	"github.com/google/uuid"
	"github.com/nyaruka/phonenumbers"
	"github.com/pkg/errors"
)

type ContactInformation struct {
	id                   uuid.UUID
	contactName          string
	emailAddress         mail.Address
	primaryPhoneNumber   string
	secondaryPhoneNumber *string
	mailingAddress       *string
}

func NewContactInformation(
	id uuid.UUID,
	contactName string,
	emailAddress string,
	primaryPhoneNumber string,
	secondaryPhoneNumber *string,
	mailingAddress *string,
) (ContactInformation, error) {
	ppn, err := parseNumber(primaryPhoneNumber)
	if err != nil {
		return ContactInformation{}, errors.WithMessagef(err, "failed to parse primary phone number for %s", contactName)
	}

	var spn *string
	if secondaryPhoneNumber != nil {
		pn, err := parseNumber(*secondaryPhoneNumber)
		if err == nil {
			spn = &pn
		}

	}

	email, err := mail.ParseAddress(emailAddress)
	if err != nil {
		return ContactInformation{}, errors.WithMessagef(err, "failed to parse email address for %s", contactName)
	}

	return ContactInformation{
		id:                   id,
		contactName:          contactName,
		emailAddress:         *email,
		primaryPhoneNumber:   ppn,
		secondaryPhoneNumber: spn,
		mailingAddress:       mailingAddress,
	}, nil
}

func parseNumber(number string) (string, error) {
	pn, err := phonenumbers.Parse(number, "US")
	if err != nil {
		return "", errors.WithStack(err)
	}

	return pn.String(), nil
}

func (ci *ContactInformation) ID() uuid.UUID {
	return ci.id
}

func (ci *ContactInformation) ContactName() string {
	return ci.contactName
}

func (ci *ContactInformation) EmailAddress() mail.Address {
	return ci.emailAddress
}

func (ci *ContactInformation) PrimaryPhoneNumber() string {
	return ci.primaryPhoneNumber
}

func (ci *ContactInformation) SetPrimaryPhoneNumber(number string) error {
	pn, err := parseNumber(number)
	if err != nil {
		return errors.WithMessagef(err, "failure to set primary number for %s", ci.contactName)
	}

	ci.primaryPhoneNumber = pn
	return nil
}

func (ci *ContactInformation) SecondaryPhoneNumber() string {
	if ci.secondaryPhoneNumber != nil {
		return *ci.secondaryPhoneNumber
	}

	return ""
}

func (ci *ContactInformation) SetSecondaryPhoneNumber(number string) error {
	pn, err := parseNumber(number)
	if err != nil {
		return errors.WithMessagef(err, "failure to set secondary number for %s", ci.contactName)
	}

	ci.secondaryPhoneNumber = &pn
	return nil
}

func (ci *ContactInformation) MailingAddress() string {
	if ci.mailingAddress != nil {
		return *ci.mailingAddress
	}

	return ""
}
