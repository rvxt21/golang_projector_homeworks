package travelagency

import "sync"

type Storage struct {
	tourM sync.Mutex
	tours []Tour
}

func (s *Storage) AddTour(t Tour) {
	s.tourM.Lock()
	defer s.tourM.Unlock()

	s.tours = append(s.tours, t)
}
