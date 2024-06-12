package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Millisecond)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAdd(t *testing.T) {
	cache := NewCache(time.Millisecond)
	cache.Add("key", []byte("val"))
	if _, ok := cache.cache["key"]; !ok {
		t.Error("key not added to cache")
	}
}

func TestGet(t *testing.T) {
	cache := NewCache(time.Millisecond)

	want := "val"

	cache.Add("key", []byte("val"))
	got, ok := cache.Get("key")
	if !ok {
		t.Errorf("key %s not found in cache", "key")
	}

	if string(got) != want {
		t.Errorf("%s does not match %s found in cache", string(got), want)
	}
}

func TestReap(t *testing.T) {
	cacheEvictionFrequency := time.Millisecond * 5
	cache := NewCache(cacheEvictionFrequency)

	cache.Add("key", []byte("val"))
	time.Sleep(cacheEvictionFrequency + time.Millisecond)

	if _, ok := cache.cache["key"]; ok {
		t.Error("Cache entries are not being reaped")
	}
}

func TestReapFailure(t *testing.T) {
	cacheEvictionFrequency := time.Millisecond * 5
	cache := NewCache(cacheEvictionFrequency)

	cache.Add("key", []byte("val"))
	time.Sleep(cacheEvictionFrequency - time.Millisecond)
	if _, ok := cache.cache["key"]; !ok {
		t.Error("Cache entries are being reaped too early")
	}
}
