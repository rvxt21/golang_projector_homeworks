package journey

import (
	"fmt"
	"hw6/publictransport"
	"reflect"
)

type Journey struct {
	RouteFrom string
	RouteTo   string
	transport []publictransport.PublicTransport
}

func NewJourney(from string, to string) *Journey {
	return &Journey{
		RouteFrom: from,
		RouteTo:   to,
		transport: []publictransport.PublicTransport{},
	}
}

func (j *Journey) StartJorney() {
	fmt.Println("Now passenger starting his journey!")
}

func (j *Journey) AddTransport(t publictransport.PublicTransport) {
	j.transport = append(j.transport, t)
}

func (j Journey) ShowAllTransport() {
	for _, transport := range j.transport {
		fmt.Printf("Type: %v\n", reflect.TypeOf(transport))
		fmt.Println("Info: ", transport)
	}
}
