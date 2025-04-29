package pokecache

import (
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	key := "testKey"
	data := []byte("testData")
	cahce := NewCache(5 * time.Second)

	cahce.Add(key, data);

	result, ok := cahce.Get(key)
	if !ok {
		t.Errorf("%s ket not found in the cache", key)
		return
	}

	if string(result) != string(data) {
		t.Errorf("stored data is not equal to expected")
		return
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Second
	const waitTine = baseTime + 5 * time.Millisecond
	const key = "testKey"
	data := []byte("testData")
	
	cache := NewCache(baseTime)
	cache.Add(key, data)
	
	_, ok := cache.Get(key)
	if !ok {
		t.Errorf("%s ket not found in the cache", key)
		return
	}

	time.Sleep(waitTine)

	_, ok = cache.Get(key)
	if ok {
		t.Errorf("%s key should not be found in the cache", key)
		return
	}
}
