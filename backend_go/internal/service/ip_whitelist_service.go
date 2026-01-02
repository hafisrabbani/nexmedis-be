package service

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"

	"github.com/hafisrabbani/technical-test-nexmedis/internal/repository"
)

type IPWhitelistService struct {
	repo *repository.IPWhitelistRepository
	rdb  *redis.Client
}

func NewIPWhitelistService(
	repo *repository.IPWhitelistRepository,
	rdb *redis.Client,
) *IPWhitelistService {
	return &IPWhitelistService{repo: repo, rdb: rdb}
}

func (s *IPWhitelistService) ReplaceAll(
	ctx context.Context,
	clientUUID string,
	ips []string,
) error {
	if err := s.repo.ReplaceAll(ctx, clientUUID, ips); err != nil {
		return err
	}

	key := fmt.Sprintf("ip_whitelist:%s", clientUUID)

	_ = s.rdb.Del(ctx, key).Err()

	if len(ips) > 0 {
		_ = s.rdb.SAdd(ctx, key, ips).Err()
	}

	return nil
}
