package travelagency

type storage interface {
	Create(t Tour)
	GetAllTours() []Tour
}

type Service struct {
	s storage
}

func NewService(s storage) *Service {
	return &Service{s: s}
}

func (s *Service) CreateTour(title string, price uint16, programm string, touristsnum uint8, nutrition Nutrition, transport Transport) {

	tour := NewTour(title, price, programm, touristsnum, nutrition, transport)

	s.s.Create(tour)
}

func (s *Service) GetAll() []Tour {
	return s.s.GetAllTours()
}
