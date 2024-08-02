package bookingservice

import idgenerator "hw15/internal/id-generator"

type BookingService struct {
	storage     storage
	idGenerator *idgenerator.IDGeneratorService
}

func NewService(s storage, id *idgenerator.IDGeneratorService) *BookingService {
	return &BookingService{storage: s,
		idGenerator: id}
}

type storage interface {
	CreateReservation(res Reservation)
	GetReservationInfo(id int) (Reservation, bool)
}

func (bs BookingService) ReserveTour(res Reservation) {
	res.ID = bs.idGenerator.GenerateID()
	bs.storage.CreateReservation(res)
}

func (bs BookingService) GetReservationInfo(id int) (Reservation, bool) {

	res, ok := bs.storage.GetReservationInfo(id)
	return res, ok
}
