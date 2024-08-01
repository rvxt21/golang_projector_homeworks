package main

import (
	"fmt"
	"hw6/publictransport/journey"
	"hw6/publictransport/passengers"
	"hw6/publictransport/publictransport"
)

func main() {
	// plane := publictransport.NewPlane()
	fmt.Println("*********************************************")
	p := passengers.NewPassenger("Ivan", "Ivanov", 20)
	p1 := passengers.NewPassenger("Petro", "Nahorniak", 36)
	p2 := passengers.NewPassenger("Bohdan", "Grubov", 40)
	plane := publictransport.NewPlane()
	fmt.Println("ADDING PASSENGERs")
	plane.TakePassengers(p)
	plane.TakePassengers(p1)
	plane.TakePassengers(p2)
	// fmt.Println("*********************************************")
	// fmt.Println(p)
	// fmt.Println(p1)
	// fmt.Println(p2)
	fmt.Println("*********************************************")
	fmt.Println(plane.Passengers)
	// fmt.Println("DELETING PASSENGERs")
	// plane.DisembarkPassengers(p)
	// fmt.Println("*********************************************")
	// fmt.Println(p)
	// fmt.Println(p1)
	// fmt.Println(p2)
	// fmt.Println(plane.Passengers)

	bus := publictransport.NewBus()
	fmt.Println("*******************ADDING PASSENGERS***********************")
	bus.TakePassengers(p2)
	bus.TakePassengers(p1)
	bus.TakePassengers(p)
	fmt.Println(bus.Passengers)
	// fmt.Println("*******************DELETING PASSENGERS***********************")
	// bus.DisembarkPassengers(p)
	// bus.DisembarkPassengers(p1)
	// bus.DisembarkPassengers(p2)
	// fmt.Println(bus.Passengers)
	// fmt.Println(p)
	// fmt.Println(p1)
	// fmt.Println(p2)

	train := publictransport.NewTrain(7)
	fmt.Println("*****************ADDING PASSENGERS********************")
	train.TakePassengers(p)
	train.TakePassengers(p1)
	train.TakePassengers(p2)
	fmt.Println(train.Passengers)
	fmt.Println(p)
	fmt.Println(p1)
	fmt.Println(p2)
	// fmt.Println("**********************DELETING PASSENGER*********************")
	// train.DisembarkPassengers(p)
	// train.DisembarkPassengers(p1)
	// train.DisembarkPassengers(p2)
	// fmt.Println(train.Passengers)
	// fmt.Println(p)
	// fmt.Println(p1)
	// fmt.Println(p2)

	journey := journey.NewJourney("Kiyv", "Warsaw")
	journey.AddTransport(train)
	journey.AddTransport(bus)
	journey.AddTransport(plane)
	train.DisembarkPassengers(p)
	train.DisembarkPassengers(p1)
	train.DisembarkPassengers(p2)

	bus.DisembarkPassengers(p)
	bus.DisembarkPassengers(p1)
	bus.DisembarkPassengers(p2)

	plane.DisembarkPassengers(p)
	// journey.ShowAllTransport()

}
