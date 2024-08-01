package main

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type InMemStorage struct {
	Oranges map[int]Orange
}

func NewInMemStorage() *InMemStorage {
	return &InMemStorage{
		Oranges: map[int]Orange{},
	}
}

var errOrangeAlreadyExists = errors.New("orange already exists")

func (m *InMemStorage) CreateOrange(o Orange) error {
	const op = "consumer-service.storage.CreateOrange"
	_, exists := m.Oranges[o.ID]
	if exists {
		log.Error().Msgf("%s: %s", op, errOrangeAlreadyExists)
		return errOrangeAlreadyExists
	}
	m.Oranges[o.ID] = o
	return nil
}

func (m *InMemStorage) GetAllOranges() []Orange {
	var oranges []Orange
	for _, o := range m.Oranges {
		oranges = append(oranges, o)
	}

	return oranges
}
