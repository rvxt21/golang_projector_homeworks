package travelagency

import (
	"errors"
	"time"
)

type Tour struct {
	ID             int
	Price          uint16
	Programm       string
	StartDate      time.Time
	EndDate        time.Time
	TouristsNumber uint8
	Nutrition      Nutrition
	TransportType  Transport
}

type Nutrition string

const (
	Breakfast             Nutrition = "breakfest"
	AllInclusive          Nutrition = "all inclusive"
	BreakfastPlusDinner   Nutrition = "breakfest and dinner"
	NotSpecifiedNutrition Nutrition = ""
)

func (n Nutrition) IsValid() error {
	switch n {
	case Breakfast, AllInclusive, BreakfastPlusDinner, NotSpecifiedNutrition:
		return nil
	default:
		return errors.New("invalid nutrition")
	}
}

type Transport string

const (
	Bus   Transport = "Bus"
	Plane Transport = "Plane"
)

func (t Transport) IsValid() error {
	switch t {
	case Bus, Plane:
		return nil
	default:
		return errors.New("invalid nutrition")
	}
}
