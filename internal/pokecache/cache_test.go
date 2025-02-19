package pokecache

import (
	"testing"
	"time"
)

func TestReapLoop(t *testing.T) {
	const baseTime = 1
	const waitTime = 2 * baseTime

	cache := NewCache(baseTime)
	cache.Add("example.com1", nil)

	_, ok := cache.Get("example.com1")
	if !ok {
		t.Errorf("Expected to find key: %s", "example.com1")
	}

	time.Sleep(waitTime * time.Second)

	_, ok = cache.Get("example.com1")
	if ok {
		t.Errorf("Expected to NOT find key: %s", "example.com1")
	}
}
