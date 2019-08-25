package cache

import (
	"testing"
	"time"
)

func TestInitCache(t *testing.T) {
	InitCache()
	Cache.Set("a", 12, time.Microsecond)
	if key, _ := Cache.Get("a"); key.(int) != 12 {
		t.Error("err", key)
	}
	time.Sleep(time.Millisecond * 3)
	if key, _ := Cache.Get("a"); key != nil {
		t.Error("err", key)
	}

}
