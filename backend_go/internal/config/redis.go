package config

import (
	"context"
	"strconv"
	"time"

	"github.com/hafisrabbani/technical-test-nexmedis/internal/shared"
	"github.com/redis/go-redis/v9"
)

func NewRedis(ctx context.Context) (*redis.Client, error) {
	db, _ := strconv.Atoi(shared.GetEnv("REDIS_DB", "0"))

	rdb := redis.NewClient(&redis.Options{
		Addr:         shared.GetEnv("REDIS_HOST", "localhost") + ":" + shared.GetEnv("REDIS_PORT", "6379"),
		Password:     shared.GetEnv("REDIS_PASSWORD", ""),
		DB:           db,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		PoolSize:     10,
		MinIdleConns: 2,
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return rdb, nil
}
