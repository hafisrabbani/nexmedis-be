package service

import (
	"context"

	"github.com/hafisrabbani/technical-test-nexmedis/internal/repository"
)

type UsageService struct {
	repo *repository.UsageRepository
}

func NewUsageService(repo *repository.UsageRepository) *UsageService {
	return &UsageService{repo: repo}
}

func (s *UsageService) Log(ctx context.Context, clientID string) {
	_ = s.repo.IncrDaily(ctx, clientID)
	_ = s.repo.IncrTop(ctx, clientID)
}

func (s *UsageService) Daily(ctx context.Context, clientID string) ([]map[string]interface{}, error) {
	return s.repo.GetDaily(ctx, clientID, 7)
}

func (s *UsageService) Top(ctx context.Context) ([]map[string]interface{}, error) {
	return s.repo.GetTop(ctx, 3)
}
