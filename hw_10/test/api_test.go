package test_tests

// import (
// 	"bytes"
// 	"encoding/json"
// 	"hw10/enteties"
// 	"hw10/resources"
// 	"hw10/storage"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/require"
// )

// func TestCreateOneTask(t *testing.T) {
// 	mockStorage := storage.NewStorage()

// 	newTask := enteties.Task{
// 		Title:       "Create API for tasks managment",
// 		Description: "Look in the Projector homework files",
// 		Priority:    enteties.HighPriority,
// 		Status:      enteties.ToDoStatus,
// 	}

// 	id, err := mockStorage.CreateOneTask(newTask)
// 	require.NoError(t, err)

// 	storedTask, exists := mockStorage.AllTasks[id]
// 	assert.True(t, exists)
// 	assert.Equal(t, newTask.Title, storedTask.Title)
// 	assert.Equal(t, newTask.Priority, storedTask.Priority)
// 	assert.Equal(t, newTask.Status, storedTask.Status)
// }

// func TestPostMethodCreatingTask(t *testing.T) {
// 	mockStorage := storage.NewStorage()
// 	mockResource := resources.TasksResourse{S: mockStorage}

// 	newTask := enteties.Task{
// 		Title:       "Create simple API",
// 		Description: "Look in the Projector homework files",
// 		Priority:    enteties.HighPriority,
// 		Status:      enteties.InProgressStatus,
// 	}

// 	jsonData, err := json.Marshal(newTask)
// 	require.NoError(t, err)
// 	req, err := http.NewRequest("POST", "/task", bytes.NewBuffer(jsonData))
// 	require.NoError(t, err)

// 	rr := httptest.NewRecorder()

// 	handler := http.HandlerFunc(mockResource.CreateTask)
// 	handler.ServeHTTP(rr, req)

// 	assert.Equal(t, http.StatusCreated, rr.Code)

// 	var createdTask enteties.Task
// 	err = json.NewDecoder(rr.Body).Decode(&createdTask)
// 	require.NoError(t, err)
// 	assert.Equal(t, newTask.Title, createdTask.Title)
// 	assert.Equal(t, newTask.Priority, createdTask.Priority)
// 	assert.Equal(t, newTask.Status, createdTask.Status)
// }
