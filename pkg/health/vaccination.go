package health

import (
	"time"
)

type Vaccination struct {
	Vaccine    Vaccine
	Vaccinated bool
	DateOfVax  time.Time
	Notes      string
}
