package userservice

import (
	"errors"
	"sync"

	"github.com/rs/zerolog/log"
)

type InMemUserStorage struct {
	m     sync.Mutex
	Users map[int]User
}

func NewInMemoryStorage() InMemUserStorage {
	return InMemUserStorage{
		Users: map[int]User{},
	}
}

func (s *InMemUserStorage) CreateUser(u User) {
	s.m.Lock()
	defer s.m.Unlock()

	s.Users[u.ID] = u
}

var ErrUserNotFound = errors.New("user not found")

func (s *InMemUserStorage) GetUserById(id int) (User, bool) {
	s.m.Lock()
	defer s.m.Unlock()

	user, ok := s.Users[id]
	if !ok {
		log.Error().Msgf("User not found")
	}
	return user, ok
}
