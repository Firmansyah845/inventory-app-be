package config

import (
	"time"

	"awesomeProjectSamb/pkg/cache"
)

var (
	CacheCfg        cache.Options
	CacheTTL        time.Duration
	CacheInvalidTTL time.Duration
	CacheApiTimeout time.Duration
	CachePrefix     string
)

func initCacheConfig() {
	CacheCfg = cache.Options{
		Addrs:        mustGetStringArray("CACHE_HOST"),
		PoolSize:     mustGetInt("CACHE_POOL_SIZE"),
		DB:           mustGetInt("CACHE_DB"),
		DialTimeout:  mustGetDurationMs("CACHE_DIAL_TIMEOUT"),
		ReadTimeout:  mustGetDurationMs("CACHE_READ_TIMEOUT"),
		WriteTimeout: mustGetDurationMs("CACHE_WRITE_TIMEOUT"),
		IdleTimeout:  mustGetDurationMs("CACHE_IDLE_TIMEOUT"),
		Password:     mustGetString("CACHE_PASSWORD"),
	}

	CacheTTL = mustGetDurationMinute("CACHE_MX_LOOKUP_TTL_MINUTES")
	CacheInvalidTTL = mustGetDurationMinute("CACHE_INVALID_MX_LOOKUP_TTL_MINUTES")
	CacheApiTimeout = mustGetDurationMinute("CACHE_API_TIMEOUT")
	CachePrefix = mustGetString("CACHE_PREFIX")

}
