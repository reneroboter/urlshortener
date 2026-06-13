package repo

import (
	"errors"
)

type ShortURLRepository struct {
	m RepositoryInterface
	r RepositoryInterface
}

func NewShortUrlRepository() *ShortURLRepository {
	return &ShortURLRepository{
		m: NewMySqlStore(),
		r: NewRedisRepo(),
	}
}

func (t *ShortURLRepository) Put(code, url string) error {
	// write to MySQL first; it's the source of truth (SSOT)
	err1 := t.m.Put(code, url)
	if err1 != nil {
		return errors.New("code already exists")
	}
	// write to memory second as a best-effort cache to reduce read latency
	_ = t.r.Put(code, url)

	return nil
}

func (t *ShortURLRepository) Get(code string) (string, error) {
	// check Redis first, it's faster
	url, err := t.r.Get(code)

	if url != "" {
		return url, nil
	}

	// TODO: only fallback on ErrNotFound, not on any error
	if err != nil {
		// check redis second if the entry is not found in memory (redis is the SSOT)
		url, err = t.m.Get(code)
	}

	if url != "" {
		// backfill Redis cache after a successful MySQL read
		_ = t.m.Put(code, url)
		return url, nil
	}

	return "", errors.New("code not found")
}
