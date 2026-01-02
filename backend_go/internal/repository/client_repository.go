package repository

import (
	"context"
	domainErr "github.com/hafisrabbani/technical-test-nexmedis/internal/model/error"
	"gorm.io/gorm"
)

type Client struct {
	ID         string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	ClientID   string `gorm:"uniqueIndex;size:64;not null"`
	Name       string `gorm:"size:255;not null"`
	Email      []byte `gorm:"not null"`
	APIKeyHash string `gorm:"not null"`
	Status     string `gorm:"type:varchar(20);default:'active'"`
}

type ClientRepository struct {
	db *gorm.DB
}

func NewClientRepository(db *gorm.DB) *ClientRepository {
	return &ClientRepository{db: db}
}

func (r *ClientRepository) Create(
	ctx context.Context,
	clientID string,
	name string,
	emailEncrypted []byte,
	apiKeyHash string,
) error {
	var count int64

	// existence check
	if err := r.db.WithContext(ctx).
		Model(&Client{}).
		Where("client_id = ?", clientID).
		Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return domainErr.ErrClientAlreadyExists
	}

	client := Client{
		ClientID:   clientID,
		Name:       name,
		Email:      emailEncrypted,
		APIKeyHash: apiKeyHash,
	}

	return r.db.WithContext(ctx).Create(&client).Error
}

func (r *ClientRepository) FindByAPIKeyHash(
	ctx context.Context,
	apiKeyHash string,
) (*Client, error) {
	var client Client

	err := r.db.WithContext(ctx).
		Where("api_key_hash = ? AND status = ?", apiKeyHash, "active").
		First(&client).Error

	if err != nil {
		return nil, err
	}

	return &client, nil
}

func (r *ClientRepository) FindByClientID(
	ctx context.Context,
	clientID string,
) (*Client, error) {
	var client Client

	err := r.db.WithContext(ctx).
		Where("client_id = ?", clientID).
		First(&client).Error

	if err != nil {
		return nil, err
	}

	return &client, nil
}

func (r *ClientRepository) GetUUIDByClientID(
	ctx context.Context,
	clientID string,
) (string, error) {
	var id string
	err := r.db.WithContext(ctx).
		Raw(`SELECT id FROM clients WHERE client_id = ?`, clientID).
		Scan(&id).Error
	return id, err
}
