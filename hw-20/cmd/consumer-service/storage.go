package main

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type InMemStorage struct {
	Oranges   map[int]Orange
	Analytics OrangesAnalytics
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
	switch {
	case o.Size >= 0 && o.Size <= 100:
		m.Analytics.Small++
	case o.Size > 100 && o.Size <= 200:
		m.Analytics.Medium++
	case o.Size > 200 && o.Size <= 300:
		m.Analytics.Big++
	}
	return nil

}

func (m *InMemStorage) GetAllOranges() []Orange {
	var oranges []Orange
	for _, o := range m.Oranges {
		oranges = append(oranges, o)
	}

	return oranges
}

func (m *InMemStorage) GetAnalytics() map[string]int {
	analytics := make(map[string]int, 3)
	analytics["small"] = m.Analytics.Small
	analytics["medium"] = m.Analytics.Medium
	analytics["big"] = m.Analytics.Big
	return analytics
}
