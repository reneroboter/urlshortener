package store

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"sync"
)

var UrlsMap = sync.Map{}

type StoreInterface interface {
	Put(url string) (string, error)
	Get(code string) (string, error)
}
type InMemoryStore struct {
	m map[string]string
}

type FileStore struct {
	m map[string]string
}

func NewInMemoryStore() *InMemoryStore {
	// todo how to ensure singleton?
	return &InMemoryStore{m: make(map[string]string)}
}

func (s *InMemoryStore) Put(url string) (string, error) {
	code := fmt.Sprintf("%x", sha1.Sum([]byte(url)))
	s.m[code] = url
	return code, nil
}

func (s *InMemoryStore) Get(code string) (string, error) {
	url, ok := s.m[code]
	if !ok {
		return "", errors.New("not found")
	}
	return url, nil
}
