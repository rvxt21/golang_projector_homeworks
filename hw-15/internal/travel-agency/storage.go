package travelagency

import (
	"sync"

	"github.com/rs/zerolog/log"
)

type InMemoryStorage struct {
	tourM sync.Mutex
	tours []Tour
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{}
}

func (s *InMemoryStorage) Create(t Tour) {
	s.tourM.Lock()
	defer s.tourM.Unlock()

	log.Info().Msg("Trying to create one task")

	if err := t.Nutrition.IsValid(); err != nil {
		log.Info().Err(err).Msgf("Not valid nutrition for trying to create %v tour.", t.ID)
		return
	}

	if err := t.TransportType.IsValid(); err != nil {
		log.Info().Err(err).Msgf("Not valid transport type for trying to create %v tour.", t.ID)
		return
	}
	s.tours = append(s.tours, t)
}

func (s *InMemoryStorage) GetAllTours() []Tour {
	s.tourM.Lock()
	defer s.tourM.Unlock()

	log.Info().Msg("Getting all tasks")

	tours := make([]Tour, len(s.tours))
	copy(tours, s.tours)

	return tours
}
