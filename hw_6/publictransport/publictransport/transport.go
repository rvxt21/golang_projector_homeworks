package publictransport

import "hw6/publictransport/passengers"

type PublicTransport interface {
	TakePassengers(*passengers.Passenger)
	DisembarkPassengers(*passengers.Passenger)
}
