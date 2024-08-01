package storage

import (
	"hw10/enteties"
	"sort"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

type InMemory struct {
	m          sync.Mutex
	lastId     int
	AllTasks   map[int]enteties.Task
	lastUserId int
	allUsers   map[int]enteties.User
}

func NewInMemory() *InMemory {
	return &InMemory{
		AllTasks: make(map[int]enteties.Task),
	}
}

func (s *InMemory) CreateOneTask(t enteties.Task) (int, error) {
	s.m.Lock()
	defer s.m.Unlock()

	log.Info().Msg("Trying to create one task")

	if err := t.Priority.IsValid(); err != nil {
		log.Info().Err(err).Msgf("Not valid priority for trying to create %d task.", t.ID)
		return 0, err
	}

	if err := t.Status.IsValid(); err != nil {
		log.Info().Err(err).Msgf("Not valid status for trying to create %d task.", t.ID)
		return 0, err
	}

	t.ID = s.lastId + 1
	s.AllTasks[t.ID] = t
	s.lastId++
	t.CreatedAt = time.Now()

	log.Info().Msgf("Created one task. ID: %d", s.lastId)
	return s.lastId, nil
}

func (s *InMemory) GetAllTasks() []enteties.Task {
	s.m.Lock()
	defer s.m.Unlock()

	log.Info().Msg("Getting all tasks")

	var tasks = make([]enteties.Task, 0, len(s.AllTasks))
	for _, task := range s.AllTasks {
		tasks = append(tasks, task)
	}

	sort.Slice(tasks, func(i, j int) bool { return tasks[i].ID < tasks[j].ID })

	return tasks
}

func (s *InMemory) DeleteTask(id int) bool {
	s.m.Lock()
	defer s.m.Unlock()

	log.Info().Msgf("Deleting task ID %d", id)

	_, ok := s.AllTasks[id]
	if !ok {
		return false
	}

	delete(s.AllTasks, id)
	return true
}

func (s *InMemory) UpdateTask(idForUpdate int, t enteties.Task) bool {
	s.m.Lock()
	defer s.m.Unlock()

	log.Info().Msgf("Updating task ID %d", idForUpdate)

	taskToUpdate, exists := s.AllTasks[idForUpdate]
	if !exists {
		log.Info().Msg("Task does not exists. Invalid to update.")
		return false
	}
	if t.Title != "" {
		taskToUpdate.Title = t.Title
	}
	if t.Description != "" {
		taskToUpdate.Description = t.Description
	}
	if t.Status != "" {
		taskToUpdate.Status = t.Status
	}
	if !t.DueDate.IsZero() {
		taskToUpdate.DueDate = t.DueDate
	}
	s.AllTasks[idForUpdate] = taskToUpdate
	return true
}

func (s *InMemory) CreateOneUser(u enteties.User) (int, error) {
	s.m.Lock()
	defer s.m.Unlock()

	log.Info().Msg("Creating one user")

	if err := u.Role.IsValid(); err != nil {
		log.Error().Err(err)
		return 0, err
	}

	u.ID = s.lastUserId + 1
	s.allUsers[u.ID] = u
	s.lastUserId++

	return s.lastUserId, nil
}
