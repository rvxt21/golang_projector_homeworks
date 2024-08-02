package bookingservice

import (
	"sync"

	"github.com/rs/zerolog/log"
)

type InMemReservationSorage struct {
	Reservations map[int]Reservation
	mu           sync.Mutex
}

func NewInMemStorage() InMemReservationSorage {
	return InMemReservationSorage{
		Reservations: map[int]Reservation{},
	}
}

func (s *InMemReservationSorage) CreateReservation(res Reservation) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.Reservations[res.ID] = res
}

func (s *InMemReservationSorage) GetReservationInfo(id int) (Reservation, bool) {
	res, ok := s.Reservations[id]
	if !ok {
		log.Info().Msgf("error to find %d reservation", id)
	}
	return res, ok
}

func (s *InMemReservationSorage) GetReservationsByUserID(userID int) []Reservation {
	s.mu.Lock()
	defer s.mu.Unlock()

	var userReservations []Reservation
	for _, reservation := range s.Reservations {
		if reservation.UserID == userID {
			userReservations = append(userReservations, reservation)
		}
	}
	return userReservations
}
