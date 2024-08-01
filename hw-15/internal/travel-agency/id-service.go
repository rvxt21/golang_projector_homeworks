package travelagency

import "sync"

type IDGeneratorService struct {
	mu     sync.Mutex
	lastId int
}

func NewIDGenerator() *IDGeneratorService {
	return &IDGeneratorService{}
}

func (ig *IDGeneratorService) GenerateID() int {
	ig.mu.Lock()
	defer ig.mu.Unlock()

	ig.lastId++
	return ig.lastId
}
