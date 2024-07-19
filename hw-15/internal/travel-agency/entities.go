package travelagency

import (
	"errors"
	"time"
)

type Tour struct {
	ID             string    `json:"id"`
	Title          string    `json:"title"`
	Price          uint16    `json:"price"`
	Programm       string    `json:"program"`
	TouristsNumber uint8     `json:"tourists_number"`
	Nutrition      Nutrition `json:"nutrition"`
	TransportType  Transport `json:"transport_type"`
}

func NewTour(title string, price uint16, programm string, touristsnum uint8, nutrition Nutrition, transport Transport) Tour {
	return Tour{
		ID:             time.Now().String(),
		Title:          title,
		Price:          price,
		Programm:       programm,
		TouristsNumber: touristsnum,
		Nutrition:      nutrition,
		TransportType:  transport,
	}
}

type Nutrition string

const (
	Breakfast             Nutrition = "Breakfast"
	AllInclusive          Nutrition = "All inclusive"
	BreakfastPlusDinner   Nutrition = "Breakfast and dinner"
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
