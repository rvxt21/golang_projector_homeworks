package publictransport

import "pubtransport/passengers"

type PublicTransport interface {
	TakePassengers(*passengers.Passenger)
	DisembarkPassengers(*passengers.Passenger)
}
