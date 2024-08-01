package publictransport

import (
	"fmt"
	"hw6/publictransport/passengers"
)

type Plane struct {
	Passengers []passengers.Passenger //[passangeer name]seat
	Seats      []string
}

func NewPlane() *Plane {
	return &Plane{
		Seats: generateSeatsSlice()}
}

func (p *Plane) TakePassengers(passenger *passengers.Passenger) {
	ticket := passengers.NewTicket("Plane", 0, 0, 0, "")
	if containsValue(p.Passengers, *passenger) {
		fmt.Printf("Error seating on the plane for the passenger %v, already on plane.\n", passenger)
	} else {

		p.seatOnFreeSeat(ticket)
		passenger.Tikets = append(passenger.Tikets, *ticket)
		p.Passengers = append(p.Passengers, *passenger)
	}

	info := fmt.Sprintf("Passenger %s %s boarded on the plane, seat %s.\n", passenger.Name, passenger.Surname, ticket.FlightSeat)
	err := passenger.AddHistory(info)
	if err != nil {
		fmt.Println("Error:", err)
	}

}

func (p *Plane) DisembarkPassengers(pas *passengers.Passenger) {
	ticket := pas.FindPlaneTicket()
	idx := -1
	if ticket.TransportType != "" {
		for i, passenger := range p.Passengers {
			for it, passTicket := range passenger.Tikets {
				if passTicket.FlightSeat == ticket.FlightSeat {
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

		p.Passengers = append(p.Passengers[:idx], p.Passengers[idx+1:]...)
	}
	info := fmt.Sprintf("Passenger %s %s unboarded from the plane, seat %s.\n", pas.Name, pas.Surname, ticket.FlightSeat)
	err := pas.AddHistory(info)
	if err != nil {
		fmt.Println("Error:", err)
	}

}

func containsValue(m []passengers.Passenger, value passengers.Passenger) bool {
	for _, val := range m {
		if value.Name == val.Name && val.Surname == value.Surname && val.Age == value.Age {
			return true
		}
	}
	return false
}

func generateSeatsSlice() []string {
	var seats []string

	columns := []string{"A", "B", "C", "D", "E", "F"}

	for row := 1; row <= 30; row++ {
		for _, column := range columns {
			seat := fmt.Sprintf("%s%d", column, row)
			seats = append(seats, seat)
		}
	}
	return seats
}

func (p *Plane) seatOnFreeSeat(t *passengers.Ticket) {
	if len(p.Seats) != 0 {
		t.FlightSeat = p.Seats[0]
		p.Seats = p.Seats[1:]
	}
}
