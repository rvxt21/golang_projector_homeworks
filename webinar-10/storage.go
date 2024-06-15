package main

import (
	"fmt"
	"sort"
)

type Storage struct {
	lastID   int
	allTrips map[int]Trip
}

func NewStorage() *Storage {
	return &Storage{
		allTrips: make(map[int]Trip),
	}
}

func (s Storage) GetAllTrips() []Trip {
	var trips = make([]Trip, 0, len(s.allTrips))

	for _, t := range s.allTrips {
		trips = append(trips, t)
	}

	sort.Slice(trips, func(i, j int) bool { return trips[i].ID < trips[j].ID })

	return trips
}

func (s Storage) CreateOneTrip(t Trip) int {
	fmt.Println("Trying to create trip")
	t.ID = s.lastID + 1
	s.allTrips[t.ID] = t
	s.lastID = t.ID
	return t.ID
	fmt.Printf("Created trip. Last ID: %v\n", s.lastID)
}
