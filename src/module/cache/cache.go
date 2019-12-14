package cache

import (
	"time"
	"conf"
	ec "middleware/cache"
)

const (
	DefaultExpiration = 3600
	DEFAULT           = time.Duration(0)
	FOREVER           = time.Duration(-1)
)

func Cache() ec.CacheStore {
	var store ec.CacheStore

	switch conf.Conf.CacheStore {
	case conf.MEMCACHED:
		store = ec.NewMemcachedStore([]string{conf.Conf.Memcached.Server}, time.Hour)
	case conf.REDIS:
		store = ec.NewRedisCache(conf.Conf.Redis.Server, conf.Conf.Redis.Pwd, DefaultExpiration)
	default:
		store = ec.NewInMemoryStore(time.Hour)
	}

	return store
}
