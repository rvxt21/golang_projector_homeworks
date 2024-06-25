package main

import (
	"sort"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

type Storage struct {
	m          sync.Mutex
	lastId     int
	allTasks   map[int]Task
	lastUserId int
	allUsers   map[int]User
}

func NewStorage() *Storage {
	return &Storage{
		allTasks: make(map[int]Task),
	}
}

func (s *Storage) CreateOneTask(t Task) (int, error) {
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
	s.allTasks[t.ID] = t
	s.lastId++
	t.CreatedAt = time.Now()

	log.Info().Msgf("Created one task. ID: %d", s.lastId)
	return s.lastId, nil
}

func (s *Storage) GetAllTasks() []Task {
	s.m.Lock()
	defer s.m.Unlock()

	log.Info().Msg("Getting all tasks")

	var tasks = make([]Task, 0, len(s.allTasks))
	for _, task := range s.allTasks {
		tasks = append(tasks, task)
	}

	sort.Slice(tasks, func(i, j int) bool { return tasks[i].ID < tasks[j].ID })

	return tasks
}

func (s *Storage) DeleteTask(id int) bool {
	s.m.Lock()
	defer s.m.Unlock()

	log.Info().Msgf("Deleting task ID %d", id)

	_, ok := s.allTasks[id]
	if !ok {
		return false
	}

	delete(s.allTasks, id)
	return true
}

func (s *Storage) UpdateTask(idForUpdate int, t Task) bool {
	s.m.Lock()
	defer s.m.Unlock()

	log.Info().Msgf("Updating task ID %d", idForUpdate)

	taskToUpdate, exists := s.allTasks[idForUpdate]
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
	s.allTasks[idForUpdate] = taskToUpdate
	return true
}

func (s *Storage) CreateOneUser(u User) (int, error) {
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
