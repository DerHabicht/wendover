package config

import (
	"time"
)

type ActivityInfo struct {
	Key   string    `yaml:"key"`
	Name  string    `yaml:"Location"`
	Start time.Time `yaml:"start_datetime"`
	End   time.Time `yaml:"end_datetime"`
}

type ActivityConfig struct {
	Activity           ActivityInfo                    `yaml:"activity"`
	RegistrationZone   RegistrationZoneColumnMapping   `yaml:"registration_zone"`
	ApplicationPackage ApplicationPackageColumnMapping `yaml:"application_package"`
}
