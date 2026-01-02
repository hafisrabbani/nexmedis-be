package service

import (
	"context"
	"errors"

	"github.com/hafisrabbani/technical-test-nexmedis/internal/repository"
	"github.com/hafisrabbani/technical-test-nexmedis/internal/shared"
)

var (
	ErrInvalidAPIKey = errors.New("invalid api key")
)

type ClientService struct {
	repo *repository.ClientRepository
}

func NewClientService(repo *repository.ClientRepository) *ClientService {
	return &ClientService{repo: repo}
}

func (s *ClientService) Register(
	ctx context.Context,
	clientID, name, email string,
) (string, error) {
	apiKey, _ := shared.GenerateApiKey()
	apiKeyHash := shared.HashApiKey(apiKey)
	emailEncrypted := shared.EncryptEmail(email)

	err := s.repo.Create(ctx, clientID, name, emailEncrypted, apiKeyHash)
	if err != nil {
		return "", err
	}

	return apiKey, nil
}

func (s *ClientService) ValidateAPIKey(
	ctx context.Context,
	apiKey string,
) (*repository.Client, error) {
	hash := shared.HashApiKey(apiKey)

	client, err := s.repo.FindByAPIKeyHash(ctx, hash)
	if err != nil {
		return nil, ErrInvalidAPIKey
	}

	return client, nil
}
