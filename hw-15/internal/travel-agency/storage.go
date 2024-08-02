package travelagency

import (
	"errors"
	"sync"

	"github.com/rs/zerolog/log"
)

type InMemoryStorage struct {
	tourM sync.Mutex
	tours map[int]Tour
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		tours: map[int]Tour{},
	}
}

func (s *InMemoryStorage) Create(t Tour) {
	s.tourM.Lock()
	defer s.tourM.Unlock()

	log.Info().Msg("Trying to create one task")
	s.tours[t.ID] = t

}

func (s *InMemoryStorage) GetAllTours() map[int]Tour {
	s.tourM.Lock()
	defer s.tourM.Unlock()

	log.Info().Msg("Getting all tasks")

	return s.tours
}

var ErrTourNotFound = errors.New("tour not found")

func (s *InMemoryStorage) GetTourByID(tourID int) (Tour, error) {
	s.tourM.Lock()
	defer s.tourM.Unlock()

	log.Info().Msgf("Getting tour ID: %d", tourID)
	tour, exists := s.tours[tourID]
	if !exists {
		return Tour{}, ErrTourNotFound
	}

	return tour, nil
}
