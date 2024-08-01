package travelagency

import (
	"errors"
)

type Tour struct {
	ID             int       `json:"id"`
	Title          string    `json:"title"`
	Price          uint64    `json:"price"`
	Programm       string    `json:"program"`
	TouristsNumber uint8     `json:"tourists_number"`
	Nutrition      Nutrition `json:"nutrition"`
	TransportType  Transport `json:"transport_type"`
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

type User struct {
	Name        string
	Email       string
	BookedTours []Tour
}

func (u *User) BookTour(tour Tour) {
	u.BookedTours = append(u.BookedTours, tour)
}
