package publictransport

import (
	"fmt"
	"pubtransport/passengers"
)

type Train struct {
	Passengers     []passengers.Passenger
	NumberOfWagons uint8
	CurrentWagon   uint8
	CurrentCoupe   uint8
	CurrentPlace   uint8
}

func NewTrain(totalWagons uint8) *Train {
	return &Train{
		NumberOfWagons: totalWagons,
		CurrentWagon:   1,
		CurrentCoupe:   1,
		CurrentPlace:   1,
	}
}

func (t *Train) TakePassengers(p *passengers.Passenger) {

	ticket := passengers.NewTicket("Train", t.CurrentWagon, t.CurrentCoupe, t.CurrentPlace, "")
	p.Tikets = append(p.Tikets, *ticket)
	t.Passengers = append(t.Passengers, *p)
	t.CurrentPlace++
	t.checkPlaces()
}

func (t *Train) checkPlaces() {
	if t.CurrentPlace > 4 {
		t.CurrentPlace = 1
		t.CurrentCoupe++
		if t.CurrentCoupe > 10 {
			t.CurrentCoupe = 1
			t.CurrentWagon++
			if t.CurrentWagon > t.NumberOfWagons {
				fmt.Println("This train is already full.")
				t.CurrentWagon = t.NumberOfWagons
				t.CurrentCoupe = 10
				t.CurrentPlace = 4
			}
		}
	}
}

func (t *Train) DisembarkPassengers(pas *passengers.Passenger) {
	ticket := pas.FindTrainTicket()
	idx := -1
	if ticket.TransportType != "" {
		for i, passenger := range t.Passengers {
			for it, passTicket := range passenger.Tikets {
				if passTicket.Wagon == ticket.Wagon && passTicket.Coupe == ticket.Coupe && passTicket.Seat == ticket.Seat {
					idx = i
					pas.Tikets = append(pas.Tikets[:it], pas.Tikets[it+1:]...)
					break
				}
			}
			if idx != -1 {
				break
			}
		}
	}

	if idx != -1 {

		t.Passengers = append(t.Passengers[:idx], t.Passengers[idx+1:]...)
	}

}
