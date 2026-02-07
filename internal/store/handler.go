package store

import (
	"errors"
	"fmt"
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

	return nil
}

func (t *TwoLayerStore) Get(code string) (string, error) {
	url, err := t.m.Get(code)

	if url != "" {
		fmt.Println("memory found")
		return url, nil
	}

	// check on code not found? more precise!
	if err != nil {
		fmt.Println("redis not found")
		url, err = t.r.Get(code)
	}

	if url != "" {
		return url, nil
	}

	return "", errors.New("code not found")
}
