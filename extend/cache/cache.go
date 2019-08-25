package cache

import (
	"time"

	gocache "github.com/patrickmn/go-cache"
)

var Cache = &gocache.Cache{}

func InitCache() {
	Cache = gocache.New(60*time.Minute, 120*time.Minute)
}
