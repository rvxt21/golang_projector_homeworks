package travelagency

import "github.com/rs/zerolog/log"

type storage interface {
	Create(t Tour)
	GetAllTours() map[int]Tour
	GetTourByID(tourID int) (Tour, error)
}

type Service struct {
	s           storage
	idGenerator *IDGeneratorService
}

func NewService(s storage, idGenerator *IDGeneratorService) *Service {
	return &Service{s: s, idGenerator: idGenerator}
}

func (s *Service) CreateTour(tour Tour) {
	if err := tour.Nutrition.IsValid(); err != nil {
		log.Info().Err(err).Msgf("Not valid nutrition for trying to create %v tour.", tour.ID)
		return
	}

	if err := tour.TransportType.IsValid(); err != nil {
		log.Info().Err(err).Msgf("Not valid transport type for trying to create %v tour.", tour.ID)
		return
	}
	tour.ID = s.idGenerator.GenerateID()

	s.s.Create(tour)
}

func (s *Service) GetAllTours() map[int]Tour {
	return s.s.GetAllTours()
}

func (s *Service) GetTourByID(tourID int) (Tour, error) {
	t, err := s.s.GetTourByID(tourID)
	if err != nil {
		return Tour{}, errTourNotFound
	}
	return t, nil
}

// func (s *Service) BookTour(id string) {
// 	for _, tour := range s.s.GetAllTours() {
// 		if tour.ID == id {
// 			s.user.BookTour(tour)
// 			break
// 		}
// 	}
// }

// func (s *Service) GetUserTours() []Tour {
// 	return s.user.BookedTours
// }