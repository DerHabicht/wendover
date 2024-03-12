package flights

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
)

type Flight struct {
	airline      Airline
	flightNumber string
}

func NewFlight(airline Airline, flightNumber string) Flight {
	return Flight{
		airline:      airline,
		flightNumber: flightNumber,
	}
}

func (flt Flight) String() string {
	return fmt.Sprintf("%s%s", flt.airline.ICAO(), flt.flightNumber)
}

type ArrivingFlight struct {
	flight  Flight
	arrival time.Time
}

func NewArrivingFlight(airline Airline, flightNumber string, arrival time.Time) ArrivingFlight {
	flight := NewFlight(airline, flightNumber)

	return ArrivingFlight{
		flight:  flight,
		arrival: arrival,
	}
}

// String shows the flight number information, encoding datetime in the Mountain Timezone. For now, only Mountain Time
// is supported.
func (af ArrivingFlight) String() string {
	denver, err := time.LoadLocation("America/Denver")
	if err != nil {
		panic(errors.WithMessage(err, "failed to look up America/Denver"))
	}
	timestr := af.arrival.In(denver).Format("2006-01-02T15:04")

	return fmt.Sprintf("Departure, %s, %s", af.flight, timestr)
}

type DepartingFlight struct {
	flight    Flight
	departure time.Time
}

func NewDepartingFlight(airline Airline, flightNumber string, departure time.Time) DepartingFlight {
	flight := NewFlight(airline, flightNumber)

	return DepartingFlight{
		flight:    flight,
		departure: departure,
	}
}

// String shows the flight number information, encoding datetime in the Mountain Timezone. For now, only Mountain Time
// is supported.
func (df DepartingFlight) String() string {
	denver, err := time.LoadLocation("America/Denver")
	if err != nil {
		panic(errors.WithMessage(err, "failed to look up America/Denver"))
	}
	timestr := df.departure.In(denver).Format("2006-01-02T15:04")

	return fmt.Sprintf("Departure, %s, %s", df.flight, timestr)
}
