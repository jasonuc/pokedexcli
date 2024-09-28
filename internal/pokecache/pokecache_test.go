package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	c := NewCache(time.Millisecond)
	if c.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddGetCache(t *testing.T) {

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{"key1", []byte("val1")},
		{"key2", []byte("val2")},
		{"key3", []byte("val3")},
		{"key4", []byte("val4")},
		{"key5", []byte("val5")},
	}

	for _, cs := range cases {
		c := NewCache(time.Millisecond)

		c.Add(cs.inputKey, cs.inputVal)
		actual, ok := c.Get(cs.inputKey)
		if !ok {
			t.Errorf("%s not found", cs.inputKey)
		}
		if string(actual) != string(cs.inputVal) {
			t.Errorf("%s doesn't match %s", string(cs.inputVal), string(actual))
		}
	}
}

func TestReap(t *testing.T) {

	interval := time.Millisecond * 10
	cache := NewCache(time.Millisecond * 10)

	keyOne := "key1"
	cache.Add(keyOne, []byte("val1"))

	time.Sleep(interval + time.Millisecond)
	if _, ok := cache.Get(keyOne); ok {
		t.Errorf("%s should not exist after %s", keyOne, interval.String())
	}
}

func TestReapFail(t *testing.T) {

	interval := time.Millisecond * 10
	cache := NewCache(time.Millisecond * 10)

	keyOne := "key1"
	cache.Add(keyOne, []byte("val1"))

	time.Sleep(interval / 2)
	if _, ok := cache.Get(keyOne); !ok {
		t.Errorf("%s should exist", keyOne)
	}
}
