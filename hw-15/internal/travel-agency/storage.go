package travelagency

import (
	"errors"
	"sync"

	"github.com/rs/zerolog/log"
)

type InMemoryStorage struct {
	tourM sync.Mutex
	Tours map[int]Tour
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		Tours: map[int]Tour{},
	}
}

func (s *InMemoryStorage) Create(t Tour) (int, error) {
	s.tourM.Lock()
	defer s.tourM.Unlock()

	log.Info().Msg("Trying to create one task")
	s.Tours[t.ID] = t
	return t.ID, nil
}

func (s *InMemoryStorage) GetAllTours() map[int]Tour {
	s.tourM.Lock()
	defer s.tourM.Unlock()

	log.Info().Msg("Getting all tasks")

	return s.Tours
}

var ErrTourNotFound = errors.New("tour not found")

func (s *InMemoryStorage) GetTourByID(tourID int) (Tour, error) {
	s.tourM.Lock()
	defer s.tourM.Unlock()

	log.Info().Msgf("Getting tour ID: %d", tourID)
	tour, exists := s.Tours[tourID]
	if !exists {
		return Tour{}, ErrTourNotFound
	}

	return tour, nil
}
