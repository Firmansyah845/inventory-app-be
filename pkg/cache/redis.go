package cache

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
)

type Client interface {
	redis.Cmdable
	Process(ctx context.Context, cmd redis.Cmder) error
	AddHook(hook redis.Hook)
	Close() error
}

type Options struct {
	Addrs                                               []string
	PoolSize, DB                                        int
	DialTimeout, ReadTimeout, WriteTimeout, IdleTimeout time.Duration
	Password                                            string
}

func NewClient(opts *Options) (Client, error) {
	if len(opts.Addrs) == 0 {
		return nil, errors.New("no address to connect")
	}

	if len(opts.Addrs) > 1 {
		return redis.NewClusterClient(opts.cluster()), nil
	}

	return redis.NewClient(opts.simple()), nil
}

func (o *Options) cluster() *redis.ClusterOptions {
	return &redis.ClusterOptions{
		Addrs:        o.Addrs,
		PoolSize:     o.PoolSize,
		DialTimeout:  o.DialTimeout,
		ReadTimeout:  o.ReadTimeout,
		WriteTimeout: o.WriteTimeout,
		IdleTimeout:  o.IdleTimeout,
		Password:     o.Password,
	}
}

func (o *Options) simple() *redis.Options {
	var addr string

	if len(o.Addrs) > 0 {
		addr = o.Addrs[0]
	}

	return &redis.Options{
		Addr:         addr,
		PoolSize:     o.PoolSize,
		DialTimeout:  o.DialTimeout,
		ReadTimeout:  o.ReadTimeout,
		WriteTimeout: o.WriteTimeout,
		IdleTimeout:  o.IdleTimeout,
		DB:           o.DB,
		Password:     o.Password,
	}
}
