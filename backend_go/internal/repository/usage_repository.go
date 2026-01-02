package repository

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"

	"github.com/redis/go-redis/v9"
)

type UsageRepository struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewUsageRepository(db *gorm.DB, rdb *redis.Client) *UsageRepository {
	return &UsageRepository{db: db, rdb: rdb}
}

func (r *UsageRepository) IncrDaily(ctx context.Context, clientID string) error {
	key := fmt.Sprintf("usage:daily:%s:%s", clientID, time.Now().Format("2006-01-02"))
	return r.rdb.Incr(ctx, key).Err()
}

func (r *UsageRepository) IncrTop(ctx context.Context, clientID string) error {
	return r.rdb.ZIncrBy(ctx, "usage:top:24h", 1, clientID).Err()
}

func (r *UsageRepository) GetDaily(ctx context.Context, clientID string, days int) ([]map[string]interface{}, error) {
	var result []map[string]interface{}

	for i := 0; i < days; i++ {
		date := time.Now().AddDate(0, 0, -i).Format("2006-01-02")
		key := fmt.Sprintf("usage:daily:%s:%s", clientID, date)

		val, err := r.rdb.Get(ctx, key).Int64()
		if err != nil && err != redis.Nil {
			return nil, err
		}

		result = append(result, map[string]interface{}{
			"date":           date,
			"total_requests": val,
		})
	}

	return result, nil
}

func (r *UsageRepository) GetTop(ctx context.Context, limit int64) ([]map[string]interface{}, error) {
	z, err := r.rdb.ZRevRangeWithScores(ctx, "usage:top:24h", 0, limit-1).Result()
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	for _, item := range z {
		result = append(result, map[string]interface{}{
			"client_id":      item.Member,
			"total_requests": int64(item.Score),
		})
	}

	return result, nil
}

func (r *UsageRepository) UpsertDailyUsage(
	ctx context.Context,
	clientID string,
	date time.Time,
	total int64,
) error {
	return r.db.WithContext(ctx).
		Exec(`
			INSERT INTO daily_usage (client_id, date, total_requests)
			VALUES (?, ?, ?)
			ON CONFLICT (client_id, date)
			DO UPDATE SET
				total_requests = EXCLUDED.total_requests,
				updated_at = NOW()
		`, clientID, date, total).
		Error
}
