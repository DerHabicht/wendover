package repositories

import (
	"github.com/google/uuid"

	"github.com/derhabicht/wendover/pkg/org"
)

type RepositoryObject interface {
	Create() error
	CreateOrUpdate() error
	Fetch(id uuid.UUID) error
	Update() error
	Delete() error
}

type ApplicationRepository interface {
}

type FlightRepository interface {
}

type HealthRepository interface {
}

type OrgRepository interface {
	CreateActivity(activity *org.Activity) error
	FetchActivity(key string) (org.Activity, error)
	UpdateActivity(activity *org.Activity) error
	DeleteActivity(activity *org.Activity) error
}

type ParticipantRepository interface {
}
