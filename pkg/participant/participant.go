package participant

import (
	"time"

	"github.com/google/uuid"
)

type Participant struct {
	id                  uuid.UUID
	capid               uint
	memberType          MemberType
	lastName            string
	firstName           string
	grade               Grade
	dob                 time.Time
	religiousPreference string
	height              uint
	weight              uint
	hairColor           HairColor
	eyeColor            EyeColor
	tShirtSize          TShirtSize
}

func (p *Participant) ID() uuid.UUID {
	return p.id
}

func (p *Participant) CAPID() uint {
	return p.capid
}

func (p *Participant) MemberType() MemberType {
	return p.memberType
}

func (p *Participant) LastName() string {
	return p.lastName
}

func (p *Participant) FirstName() string {
	return p.firstName
}

func (p *Participant) Grade() Grade {
	return p.grade
}

func (p *Participant) DOB() time.Time {
	return p.dob
}

func (p *Participant) ReligiousPreference() string {
	return p.religiousPreference
}

func (p *Participant) Height() uint {
	return p.height
}

func (p *Participant) Weight() uint {
	return p.weight
}

func (p *Participant) HairColor() HairColor {
	return p.hairColor
}

func (p *Participant) EyeColor() EyeColor {
	return p.eyeColor
}

func (p *Participant) TShirtSize() TShirtSize {
	return p.tShirtSize
}
