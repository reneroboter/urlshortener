package store

import (
	"errors"
	"sync"
)

type InMemoryStore struct {
	mu sync.RWMutex
	m  map[string]string
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		m: make(map[string]string),
	}
}
func (s *InMemoryStore) Put(code, url string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.m[code]; ok {
		return errors.New("code already exists")
	}

	s.m[code] = url
	return nil
}

func (s *InMemoryStore) Get(code string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	url, ok := s.m[code]
	if !ok {
		return "", errors.New("code not found")
	}
	return url, nil
}
