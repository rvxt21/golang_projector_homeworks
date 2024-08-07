package test

import (
	travelagency "hw15/internal/travel-agency"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateOneTour(t *testing.T) {
	mockStorage := travelagency.NewInMemoryStorage()

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
