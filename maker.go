package dbredis

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var ErrConfigNotFound = errors.New("redis config not found")

type Maker interface {
	MakeRedis(name string) (redis.UniversalClient, error)
}

type RedisMaker struct {
	Channels Channels
}

func (m RedisMaker) MakeRedis(name string) (redis.UniversalClient, error) {
	if cfg, ok := m.Channels[name]; ok {
		client := universalClient(cfg)

		if cfg.Ping {
			ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
			defer cancel()

			if res, err1 := client.Ping(ctx).Result(); err1 != nil {
				return nil, fmt.Errorf("could not check Redis server. %w", err1)
			} else if res != "PONG" {
				return nil, fmt.Errorf("could not check Redis server by configuration `%s`", name)
			}
		}

		return client, nil
	}

	return nil, fmt.Errorf("%w: `%s`", ErrConfigNotFound, name)
}

func universalClient(cfg Config) redis.UniversalClient {
	return redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:                 cfg.Addrs,
		ClientName:            cfg.ClientName,
		DB:                    cfg.DB,
		Username:              cfg.Username,
		Password:              cfg.Password,
		SentinelUsername:      cfg.SentinelUsername,
		SentinelPassword:      cfg.SentinelPassword,
		MaxRetries:            cfg.MaxRetries,
		MinRetryBackoff:       cfg.MaxRetryBackoff,
		MaxRetryBackoff:       cfg.MaxRetryBackoff,
		DialTimeout:           cfg.DialTimeout,
		ReadTimeout:           cfg.ReadTimeout,
		WriteTimeout:          cfg.WriteTimeout,
		ContextTimeoutEnabled: cfg.ContextTimeoutEnabled,
		PoolFIFO:              cfg.PoolFIFO,
		PoolSize:              cfg.PoolSize,
		PoolTimeout:           cfg.PoolTimeout,
		MinIdleConns:          cfg.MinIdleConns,
		MaxIdleConns:          cfg.MaxIdleConns,
		ConnMaxIdleTime:       cfg.ConnMaxIdleTime,
		ConnMaxLifetime:       cfg.ConnMaxLifetime,
		MaxRedirects:          cfg.MaxRedirects,
		ReadOnly:              cfg.ReadOnly,
		RouteByLatency:        cfg.RouteByLatency,
		RouteRandomly:         cfg.RouteRandomly,
		MasterName:            cfg.MasterName,
	})
}
