package passengers

type Ticket struct {
	TransportType string
	Wagon         uint8
	Coupe         uint8
	Seat          uint8
	FlightSeat    string
}

func NewTicket(tr string, wagon uint8, coupe uint8, seat uint8, fseat string) *Ticket {
	return &Ticket{
		TransportType: tr,
		Wagon:         wagon,
		Coupe:         coupe,
		Seat:          seat,
		FlightSeat:    fseat,
	}
}

type Passenger struct {
	Name    string
	Surname string
	Age     uint8
	Tikets  []Ticket
}

func NewPassenger(name string, surname string, age uint8) *Passenger {
	return &Passenger{
		Name:    name,
		Surname: surname,
		Age:     age,
	}
}

func (p Passenger) FindPlaneTicket() Ticket {
	for _, ticket := range p.Tikets {
		if ticket.TransportType == "Plane" {
			return ticket
		}
	}
	return Ticket{}
}

func (p Passenger) FindBusTicket() Ticket {
	for _, ticket := range p.Tikets {
		if ticket.TransportType == "Bus" {
			return ticket
		}
	}
	return Ticket{}
}

func (p Passenger) FindTrainTicket() Ticket {
	for _, ticket := range p.Tikets {
		if ticket.TransportType == "Train" {
			return ticket
		}
	}
	return Ticket{}
}
