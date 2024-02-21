package flights

import (
	"strings"

	"github.com/pkg/errors"
)

// Airline is an enum representing the air carrier encoded in FlightInformation. For now, only airlines that service
// KSLC are listed.
type Airline int

const (
	AeroMexico = iota
	AirCanada
	Alaska
	AmericanAirlines
	Delta
	Frontier
	HawaiianAirlines
	JetBlue
	KLM
	Spirit
	Southwest
	SunCountryAirlines
	United
)

func ParseAirline(s string) (Airline, error) {
	switch strings.ToUpper(s) {
	case "AM":
		fallthrough
	case "AMX":
		fallthrough
	case "AEROMEXICO":
		return AeroMexico, nil
	case "AC":
		fallthrough
	case "ACA":
		fallthrough
	case "AIR CANADA":
		return AirCanada, nil
	case "AS":
		fallthrough
	case "ASA":
		fallthrough
	case "ALASKA":
		return Alaska, nil
	case "AA":
		fallthrough
	case "AAL":
		fallthrough
	case "AMERICAN AIRLINES":
		return AmericanAirlines, nil
	case "DL":
		fallthrough
	case "DAL":
		fallthrough
	case "DELTA":
		return Delta, nil
	case "F9":
		fallthrough
	case "FFT":
		fallthrough
	case "FRONTIER":
		return Frontier, nil
	case "HA":
		fallthrough
	case "HAL":
		fallthrough
	case "HAWAIIAN AIRLINES":
		return HawaiianAirlines, nil
	case "B6":
		fallthrough
	case "JBU":
		fallthrough
	case "JETBLUE":
		return JetBlue, nil
	case "KL":
		fallthrough
	case "KLM":
		return KLM, nil
	case "NK":
		fallthrough
	case "NKS":
		fallthrough
	case "SPIRIT":
		return Spirit, nil
	case "WN":
		fallthrough
	case "SWA":
		fallthrough
	case "SOUTHWEST":
		return Southwest, nil
	case "SY":
		fallthrough
	case "SCX":
		fallthrough
	case "SUN COUNTRY AIRLINES":
		return SunCountryAirlines, nil
	case "UA":
		fallthrough
	case "UAL":
		fallthrough
	case "UNITED":
		return United, nil
	default:
		return -1, errors.Errorf("unrecognized airline: %s", s)
	}
}

func (a Airline) IATA() string {
	switch a {
	case AeroMexico:
		return "AM"
	case AirCanada:
		return "AC"
	case Alaska:
		return "AS"
	case AmericanAirlines:
		return "AA"
	case Delta:
		return "DL"
	case Frontier:
		return "FFT"
	case HawaiianAirlines:
		return "HA"
	case JetBlue:
		return "B6"
	case KLM:
		return "KLM"
	case Spirit:
		return "NK"
	case Southwest:
		return "WN"
	case SunCountryAirlines:
		return "SY"
	case United:
		return "UA"
	default:
		panic(errors.Errorf("invalid value for Airline: %d", a))
	}
}

func (a Airline) ICAO() string {
	switch a {
	case AeroMexico:
		return "AMX"
	case AirCanada:
		return "ACA"
	case Alaska:
		return "ASA"
	case AmericanAirlines:
		return "AAL"
	case Delta:
		return "DAL"
	case Frontier:
		return "FFT"
	case HawaiianAirlines:
		return "HAL"
	case JetBlue:
		return "B6"
	case KLM:
		return "KLM"
	case Spirit:
		return "NKS"
	case Southwest:
		return "SWA"
	case SunCountryAirlines:
		return "SCX"
	case United:
		return "UAL"
	default:
		panic(errors.Errorf("invalid value for Airline: %d", a))
	}
}

func (a Airline) String() string {
	switch a {
	case AeroMexico:
		return "AeroMexico"
	case AirCanada:
		return "Air Canada"
	case Alaska:
		return "Alaska"
	case AmericanAirlines:
		return "American Airlines"
	case Delta:
		return "Delta"
	case Frontier:
		return "Frontier"
	case HawaiianAirlines:
		return "Hawaiian Airlines"
	case JetBlue:
		return "JetBlue"
	case KLM:
		return "KLM"
	case Spirit:
		return "Spirit"
	case Southwest:
		return "Southwest"
	case SunCountryAirlines:
		return "Sun Country Airlines"
	case United:
		return "United"
	default:
		panic(errors.Errorf("invalid value for Airline: %d", a))
	}
}
