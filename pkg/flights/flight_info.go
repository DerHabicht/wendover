package flights

type FlightInformation struct {
	inbound  ArrivingFlight
	outbound DepartingFlight
}

func NewFlightInformation(inbound ArrivingFlight, outbound DepartingFlight) FlightInformation {
	return FlightInformation{
		inbound:  inbound,
		outbound: outbound,
	}
}

func (fi FlightInformation) Inbound() ArrivingFlight {
	return fi.inbound
}

func (fi FlightInformation) Outbound() DepartingFlight {
	return fi.outbound
}
