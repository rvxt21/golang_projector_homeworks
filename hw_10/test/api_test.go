package test_tests

import (
	"hw10/enteties"
	"hw10/storage"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateOneTask(t *testing.T) {
	mockStorage := storage.NewStorage()

	newTask := enteties.Task{
		Title:       "Create API for tasks managment",
		Description: "Look in the Projector homework files",
		Priority:    enteties.HighPriority,
		Status:      enteties.ToDoStatus,
	}

	id, err := mockStorage.CreateOneTask(newTask)
	require.NoError(t, err)

	storedTask, exists := mockStorage.AllTasks[id]
	assert.True(t, exists)
	assert.Equal(t, newTask.Title, storedTask.Title)
	assert.Equal(t, newTask.Priority, storedTask.Priority)
	assert.Equal(t, newTask.Status, storedTask.Status)
}
