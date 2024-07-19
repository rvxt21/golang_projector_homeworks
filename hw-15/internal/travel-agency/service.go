package travelagency

type storage interface {
	Create(t Tour)
	GetAllTours() []Tour
}

type Service struct {
	s    storage
	user *User
}

func NewService(s storage, user *User) *Service {
	return &Service{s: s, user: user}
}

func (s *Service) CreateTour(title string, price uint16, programm string, touristsnum uint8, nutrition Nutrition, transport Transport) {

	tour := NewTour(title, price, programm, touristsnum, nutrition, transport)

	s.s.Create(tour)
}

func (s *Service) GetAll() []Tour {
	return s.s.GetAllTours()
}

func (s *Service) BookTour(id string) {
	for _, tour := range s.s.GetAllTours() {
		if tour.ID == id {
			s.user.BookTour(tour)
			break
		}
	}
}

func (s *Service) GetUserTours() []Tour {
	return s.user.BookedTours
}
