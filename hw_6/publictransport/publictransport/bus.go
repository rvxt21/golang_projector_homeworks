package publictransport

import (
	"fmt"
	"pubtransport/passengers"
)

type Bus struct {
	Passengers []passengers.Passenger
	SeatsTaken uint8
	TotalSeats uint8
}

func NewBus() *Bus {
	return &Bus{
		SeatsTaken: 0,
		TotalSeats: 3,
	}
}

func (b Bus) PrintBusPassengers() {
	totalPassangeers := len(b.Passengers)
	fmt.Printf("In bus %d passangeers, seats are taken: %d.\n", totalPassangeers, b.SeatsTaken)
	fmt.Println(b.Passengers)

}

func (b *Bus) TakePassengers(p *passengers.Passenger) {
	if b.isFreeSeats() {
		b.SeatsTaken += 1
		ticket := passengers.NewTicket("Bus", 0, 0, b.SeatsTaken, "")
		p.Tikets = append(p.Tikets, *ticket)
		fmt.Printf("Passenger %s %s took the seat in bus.\n", p.Name, p.Surname)
		b.Passengers = append(b.Passengers, *p)
	} else {
		fmt.Println("There is no free standing places.Bus is full")
	}

}

func (b Bus) isFreeSeats() bool {
	return b.SeatsTaken < b.TotalSeats
}

func (b *Bus) DisembarkPassengers(pas *passengers.Passenger) {
	ticket := pas.FindBusTicket()
	idx := -1
	if ticket.TransportType != "" {
		for i, passenger := range b.Passengers {
			for it, passTicket := range passenger.Tikets {
				if passTicket.Seat == ticket.Seat {
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

		b.Passengers = append(b.Passengers[:idx], b.Passengers[idx+1:]...)
		b.SeatsTaken -= 1
	}

}
