package koreainvestment

import (
	"testing"
	"time"
)

func TestInvalidateDefault(t *testing.T) {
	cache := NewTokenCache()
	if !cache.isInvalidated() {
		t.Error("default is not invalidated")
	}
}

func TestSetToken(t *testing.T) {
	cache := NewTokenCache()
	expected := "foo"
	cache.Set(expected)
	if cache.isInvalidated() {
		t.Error("unexpectly invalidate is true")
	}

	val, err := cache.Get()
	if err != nil {
		t.Error("get error is not nil")
	}
	if val != expected {
		t.Error("get value is wrong")
	}
}

func TestSetTokenInvalidate(t *testing.T) {
	cache := &TokenCache{data: "", expireDate: time.Now(),
		invalidateInterval: time.Second * 1}
	expected := "foo"
	cache.Set(expected)
	time.Sleep(time.Second * 3)
	if !cache.isInvalidated() {
		t.Error("unexpectly invalidate is true")
	}

	val, err := cache.Get()
	if err == nil {
		t.Error("get error is not nil")
	}
	if val == expected {
		t.Error("get value is wrong")
	}
}
