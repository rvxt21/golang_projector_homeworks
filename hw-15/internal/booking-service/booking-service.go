package bookingservice

import (
	idgenerator "hw15/internal/id-generator"
	travelagency "hw15/internal/travel-agency"
	userservice "hw15/internal/user-service"

	"github.com/rs/zerolog/log"
)

type BookingService struct {
	storage      storage
	userService  *userservice.UserService
	toursService *travelagency.Service
	idGenerator  *idgenerator.IDGeneratorService
}

func NewService(s storage, id *idgenerator.IDGeneratorService, us *userservice.UserService, ts *travelagency.Service) *BookingService {
	return &BookingService{storage: s,
		userService:  us,
		toursService: ts,
		idGenerator:  id}
}

type storage interface {
	CreateReservation(res Reservation)
	GetReservationInfo(id int) (Reservation, bool)
}

func (bs BookingService) ReserveTour(res Reservation) error {

	_, ok := bs.userService.GetUserById(res.UserID)
	if !ok {
		log.Error().Msgf("wrong user ID in reservation")
		return userservice.ErrUserNotFound
	}

	_, err := bs.toursService.GetTourByID(res.TourID)
	if err != nil {
		log.Error().Msgf("wrong tour ID in reservation")
		return err
	}

	res.ID = bs.idGenerator.GenerateID()
	bs.storage.CreateReservation(res)
	return nil
}

func (bs BookingService) GetReservationInfo(id int) (Reservation, bool) {

	res, ok := bs.storage.GetReservationInfo(id)
	return res, ok
}
