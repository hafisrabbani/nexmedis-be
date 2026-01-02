package service

import (
	"context"
	"strconv"

	"github.com/hafisrabbani/technical-test-nexmedis/internal/repository"
	"github.com/hafisrabbani/technical-test-nexmedis/internal/shared"

	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

type UsageBatchWorker struct {
	repo       *repository.UsageRepository
	clientRepo *repository.ClientRepository
	rdb        *redis.Client
}

func NewUsageBatchWorker(
	repo *repository.UsageRepository,
	clientRepo *repository.ClientRepository,
	rdb *redis.Client,
) *UsageBatchWorker {
	return &UsageBatchWorker{
		repo:       repo,
		clientRepo: clientRepo,
		rdb:        rdb,
	}
}

func (w *UsageBatchWorker) Start(ctx context.Context) {
	interval, _ := strconv.Atoi(
		shared.GetEnv("INTERVAL_INSERT_DATA", "30"),
	)
	ticker := time.NewTicker(time.Duration(interval) * time.Second)

	go func() {
		for {
			select {
			case <-ticker.C:
				w.flush(ctx)
			case <-ctx.Done():
				ticker.Stop()
				return
			}
		}
	}()
}

func (w *UsageBatchWorker) flush(ctx context.Context) {
	keys, err := w.rdb.Keys(ctx, "usage:daily:*").Result()
	if err != nil {
		return
	}

	for _, key := range keys {
		parts := strings.Split(key, ":")
		if len(parts) != 4 {
			continue
		}

		clientID := parts[2] // string client_id
		date, err := time.Parse("2006-01-02", parts[3])
		if err != nil {
			continue
		}

		count, err := w.rdb.Get(ctx, key).Int64()
		if err != nil {
			continue
		}

		clientUUID, err := w.clientRepo.GetUUIDByClientID(ctx, clientID)
		if err != nil {
			continue
		}

		_ = w.repo.UpsertDailyUsage(
			ctx,
			clientUUID,
			date,
			count,
		)
	}
}
