package cache

import (
	"context"

	"awesomeProjectSamb/internal/config"
	"awesomeProjectSamb/pkg/cache"
	apmgoredis "go.elastic.co/apm/module/apmgoredisv8/v2"
)

func InitCache() cache.Client {
	client, err := cache.NewClient(&cache.Options{
		Addrs:        config.CacheCfg.Addrs,
		PoolSize:     config.CacheCfg.PoolSize,
		DB:           config.CacheCfg.DB,
		DialTimeout:  config.CacheCfg.DialTimeout,
		ReadTimeout:  config.CacheCfg.ReadTimeout,
		WriteTimeout: config.CacheCfg.WriteTimeout,
		IdleTimeout:  config.CacheCfg.IdleTimeout,
		Password:     config.CacheCfg.Password,
	})

	client.AddHook(apmgoredis.NewHook())

	if err != nil {
		panic("[InitCache.NewClient] Could not initiate a new CacheClient : " + err.Error())
	}

	if err = client.Ping(context.Background()).Err(); err != nil {
		panic("[InitCache.NewClient.ping] Failed to connect to cache client : " + err.Error())
	}

	return client
}
