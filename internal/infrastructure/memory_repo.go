package infrastructure

import (
	"errors"
	"sync"
)

type InMemoryRepository struct {
	mu sync.RWMutex
	m  map[string]string
}

func NewInMemoryStore() *InMemoryRepository {
	return &InMemoryRepository{
		m: make(map[string]string),
	}
}
func (s *InMemoryRepository) Put(code, url string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.m[code] = url
	return nil
}

func (s *InMemoryRepository) Get(code string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	url, ok := s.m[code]
	if !ok {
		return "", errors.New("code not found")
	}
	return url, nil
}
