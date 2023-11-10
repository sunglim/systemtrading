package koreainvestment

import (
	"errors"
	"time"
)

func NewTokenCache() *TokenCache {
	return &TokenCache{data: "", expireDate: time.Now(),
		invalidateInterval: time.Minute * 5}
}

type TokenCache struct {
	data               string
	expireDate         time.Time
	invalidateInterval time.Duration
}

func (t *TokenCache) isInvalidated() bool {
	if t.expireDate.After(time.Now()) {
		return false
	}

	t.invalidate()
	return true
}

func (t *TokenCache) Get() (string, error) {
	if t.isInvalidated() {
		return "", errors.New("invalid token")
	}

	return t.data, nil
}

func (t *TokenCache) Set(data string) {
	t.data = data
	t.expireDate = time.Now().Add(t.invalidateInterval)
}

func (t *TokenCache) invalidate() {
	t.data = ""
}
