package store

import (
	"errors"
)

type TwoLayerStore struct {
	m GeneralStoreInterface
	r GeneralStoreInterface
}

func NewTwoLayerStore() *TwoLayerStore {
	return &TwoLayerStore{
		m: NewInMemoryStore(),
		r: NewRedisStore(),
	}
}

func (t *TwoLayerStore) Put(code, url string) error {
	// write to redis first; it's the source of truth (SSOT)
	err1 := t.r.Put(code, url)

	if err1 != nil {
		return errors.New("code already exists")
	}
	// write to memory second as a best-effort cache to reduce read latency
	_ = t.m.Put(code, url)

	return nil
}

func (t *TwoLayerStore) Get(code string) (string, error) {
	// check memory first, it's faster
	url, err := t.m.Get(code)

	if url != "" {
		return url, nil
	}

	// TODO: only fallback on ErrNotFound, not on any error
	if err != nil {
		// check redis second if the entry is not found in memory (redis is the SSOT)
		url, err = t.r.Get(code)
	}

	if url != "" {
		// backfill memory cache after a successful redis read
		_ = t.m.Put(code, url)
		return url, nil
	}

	return "", errors.New("code not found")
}
