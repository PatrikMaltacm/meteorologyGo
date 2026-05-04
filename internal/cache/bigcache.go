package cache

import (
	"context"
	"time"

	"github.com/allegro/bigcache/v3"
)

func NewBigCache() (*bigcache.BigCache, error) {
	config := bigcache.Config{
		Shards:             1024,
		LifeWindow:         10 * time.Minute,
		CleanWindow:        5 * time.Minute,
		MaxEntriesInWindow: 1000 * 10 * 60,
		MaxEntrySize:       500,
		Verbose:            false,
		HardMaxCacheSize:   512, // MB
	}

	return bigcache.New(context.Background(), config)
}
