package bookingservice

import (
	"time"
)

type Reservation struct {
	ID        int
	UserID    int
	TourID    int
	CreatedAt time.Time
}
