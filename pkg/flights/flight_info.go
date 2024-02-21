package flights

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
)

type FlightNumber struct {
	datetime     time.Time
	airline      Airline
	flightNumber uint
}

func NewFlightNumber(datetime time.Time, airline Airline, flightNumber uint) FlightNumber {
	return FlightNumber{
		datetime:     datetime,
		airline:      airline,
		flightNumber: flightNumber,
	}
}

// String shows the flight number information, encoding datetime in the Mountain Timezone. For now, only Mountain Time
// is supported.
func (fn FlightNumber) String() string {
	denver, err := time.LoadLocation("America/Denver")
	if err != nil {
		panic(errors.WithMessage(err, "failed to look up America/Denver"))
	}
	timestr := fn.datetime.In(denver).Format("2006-01-02T15:04")
	return fmt.Sprintf("%s%d-%s", fn.airline.ICAO(), fn.flightNumber, timestr)
}

type FlightInformation struct {
	inbound  FlightNumber
	outbound FlightNumber
}
