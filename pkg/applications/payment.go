package applications

import (
	"time"

	"github.com/google/uuid"
)

type PaymentCard struct {
	id         uuid.UUID
	Last4      uint
	NameOnCard string
	Expiration time.Time
}

type Transaction struct {
	id            uuid.UUID
	TransactionID string
	Successful    bool
	SettledStatus string
	SettledAmount uint
	CreditVoid    string
	CreditAmount  uint
	Card          PaymentCard
	CAPIDS        []uint
}

type Payment struct {
	ParticipantID uuid.UUID
	TransactionID *uuid.UUID
	PaymentAmount uint
	Scholarship   bool
}
