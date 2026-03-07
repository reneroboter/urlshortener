package infrastructure

import (
	"context"
	"fmt"
	"log/slog"
	"sync"

	"github.com/redis/go-redis/v9"
	redisclient "github.com/reneroboter/urlshortener/pkg/redis"
)

type RedisRepo struct {
	mu sync.RWMutex
	r  *redis.Client
}

func NewRedisRepo() *RedisRepo {
	return &RedisRepo{
		r: redisclient.NewRedisClient(),
	}
}

func (s *RedisRepo) Put(code, url string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	ctx := context.Background()

	_, err := s.r.SetNX(ctx, code, url, 0).Result()
	if err != nil {
		slog.Warn("[REDIS] Requested code is already stored", "code", code)
		return fmt.Errorf("[REDIS] Requested code is already stored: %w", ErrAlreadyExists)
	}
	return nil
}

func (s *RedisRepo) Get(code string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	ctx := context.Background()
	url, err := s.r.Get(ctx, code).Result()
	if err != nil {
		slog.Warn("[REDIS] Requested code could not be found!")
		return "", fmt.Errorf("[REDIS] Requested code could not be found: %w", ErrNotFound)
	}
	return url, nil
}
