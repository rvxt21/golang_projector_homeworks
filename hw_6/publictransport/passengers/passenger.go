package passengers

import (
	"fmt"
	"os"
	"time"
)

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

func (p *Passenger) AddHistory(info string) error {
	filename := p.Name + "_" + p.Surname + ".txt"
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	currentTime := time.Now()
	history := fmt.Sprintf("[%s] %s\n", currentTime.Format("2006-01-02 15:04:05"), info)
	_, err = file.WriteString(history)
	if err != nil {
		return err
	}

	return nil
}
