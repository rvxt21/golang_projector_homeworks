package test

import (
	bookingservice "hw15/internal/booking-service"
	idgenerator "hw15/internal/id-generator"
	travelagency "hw15/internal/travel-agency"
	userservice "hw15/internal/user-service"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	mockStorage        = travelagency.NewInMemoryStorage()
	mockStorageUsers   = userservice.NewInMemoryStorage()
	mockStorageBooking = bookingservice.NewInMemStorage()
	mockIdGenerator    = idgenerator.NewIDGenerator()
	mockUserService    = userservice.NewUserService(&mockStorageUsers)
	mockToursService   = travelagency.NewService(mockStorage, mockIdGenerator)
	mockBookingService = bookingservice.NewService(&mockStorageBooking, mockIdGenerator, mockUserService, mockToursService)
)

func TestCreateOneTour(t *testing.T) {

	newTour := travelagency.Tour{
		Title:          "Подорож до Таїланду без візи",
		Price:          51790,
		Programm:       "Пхукет —  тут можна загубитися серед пишної зелені джунглів, насолоджуючись ароматами екзотичних квітів. А ввечері, коли небо палає всіма відтінками помаранчевого та рожевого, здається, що навіть зірки танцюють у такт серцебиття цього дивовижного острова",
		TouristsNumber: 3,
		Nutrition:      "Breakfast and dinner",
		TransportType:  "Plane",
	}

	id, err := mockStorage.Create(newTour)
	require.NoError(t, err)

	storedTask, exists := mockStorage.Tours[id]
	assert.True(t, exists)
	assert.Equal(t, newTour.ID, storedTask.ID)
	assert.Equal(t, newTour.Title, storedTask.Title)
	assert.Equal(t, newTour.Programm, storedTask.Programm)
	assert.Equal(t, newTour.TouristsNumber, storedTask.TouristsNumber)
	assert.Equal(t, newTour.Price, storedTask.Price)
	assert.Equal(t, newTour.Nutrition, storedTask.Nutrition)
	assert.Equal(t, newTour.TransportType, storedTask.TransportType)
}

func TestGetTour(t *testing.T) {
	tour1 := travelagency.Tour{ID: 1, Title: "Test Tour", Price: 100000, Programm: "some programm", TouristsNumber: 2, Nutrition: "Breakfast", TransportType: "Plane"}
	tour2 := travelagency.Tour{ID: 2, Title: "2 Test Tour", Price: 189001, Programm: "some programm", TouristsNumber: 2, Nutrition: "All inclusive", TransportType: "Plane"}

	mockStorage.Tours[tour1.ID] = tour1
	mockStorage.Tours[tour2.ID] = tour2

	t.Run("existing tour", func(t *testing.T) {
		gotTour, err := mockStorage.GetTourByID(1)
		require.NoError(t, err)
		assert.Equal(t, tour1, gotTour)
	})

	t.Run("non-existing tour", func(t *testing.T) {
		_, err := mockStorage.GetTourByID(3)
		assert.ErrorIs(t, err, travelagency.ErrTourNotFound)
	})
}

func TestReservation(t *testing.T) {

	user1 := userservice.User{ID: 1, Name: "Anastasiia", Surname: "Koriahina", Email: "anastasiia@email.com"}
	user2 := userservice.User{ID: 2, Name: "Anna", Surname: "Cherednichenko", Email: "annache@email.com"}
	user3 := userservice.User{ID: 2, Name: "Ivan", Surname: "Hnatko", Email: "ivanhnatko@email.com"}
	mockStorageUsers.Users[user1.ID] = user1
	mockStorageUsers.Users[user2.ID] = user2
	mockStorageUsers.Users[user3.ID] = user3

	var res bookingservice.Reservation
	res.TourID = 1
	res.UserID = 2
	res.CreatedAt = time.Now()
	t.Run("existing tour+user", func(t *testing.T) {
		err := mockBookingService.ReserveTour(res)
		require.NoError(t, err)
	})

	var res2 bookingservice.Reservation
	res2.TourID = 1
	res2.UserID = 10
	res2.CreatedAt = time.Now()
	t.Run("existing tour& not existing user", func(t *testing.T) {
		err := mockBookingService.ReserveTour(res2)
		require.Error(t, err)
		require.Equal(t, userservice.ErrUserNotFound, err)
	})

}
