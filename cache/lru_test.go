package cache

import (
	"testing"
)

func TestLRU(t *testing.T) {
	capacity := 2
	c := NewLRUCache(capacity)

	c.Set(1, "Tom")
	c.Set(2, "Mike")
	c.Set(1, "Jack")
	if c.Size() != capacity {
		t.Errorf("LRU Cache size error %d\n", c.Size())
	}

	if v, ok := c.Get(1); !ok || v.(string) != "Jack" {
		t.Errorf("LRU Cache get(1) error %s\n", v)
	}

	if v, ok := c.Get(2); !ok || v.(string) != "Mike" {
		t.Errorf("LRU Cache get(2) error %s\n", v)
	}

	c.Purge()
	if c.Size() != 0 {
		t.Errorf("LRU Cache purge error")
	}
}
