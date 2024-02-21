package health

import (
	"github.com/google/uuid"
)

type History struct {
	id               uuid.UUID
	FoodAllergies    string
	OtherAllergies   string
	HealthConditions []Condition
	Details          string
	Vaccines         map[Vaccine]Vaccination
}
