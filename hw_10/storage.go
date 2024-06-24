package main

import (
	"sort"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

type Storage struct {
	m        sync.Mutex
	lastId   int
	allTasks map[int]Task
}

func NewStorage() *Storage {
	return &Storage{
		allTasks: make(map[int]Task),
	}
}

func (s *Storage) CreateOneTask(t Task) int {
	s.m.Lock()
	defer s.m.Unlock()

	log.Info().Msg("Trying to create one task")

	t.ID = s.lastId + 1
	s.allTasks[t.ID] = t
	s.lastId++
	t.CreatedAt = time.Now()

	log.Info().Msgf("Created one task. ID: %d", s.lastId)
	return s.lastId
}

func (s *Storage) GetAllTasks() []Task {
	s.m.Lock()
	defer s.m.Unlock()

	var tasks = make([]Task, 0, len(s.allTasks))
	for _, task := range s.allTasks {
		tasks = append(tasks, task)
	}

	sort.Slice(tasks, func(i, j int) bool { return tasks[i].ID < tasks[j].ID })

	return tasks
}
