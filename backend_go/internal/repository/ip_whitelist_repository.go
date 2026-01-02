package repository

import (
	"context"

	"gorm.io/gorm"
)

type IPWhitelistRepository struct {
	db *gorm.DB
}

func NewIPWhitelistRepository(db *gorm.DB) *IPWhitelistRepository {
	return &IPWhitelistRepository{db: db}
}

func (r *IPWhitelistRepository) ReplaceAll(
	ctx context.Context,
	clientUUID string,
	ips []string,
) error {
	tx := r.db.WithContext(ctx).Begin()

	// 1. delete existing
	if err := tx.Exec(
		`DELETE FROM client_ip_whitelists WHERE client_id = ?`,
		clientUUID,
	).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 2. insert new
	for _, ip := range ips {
		if ip == "" {
			continue
		}

		if err := tx.Exec(`
			INSERT INTO client_ip_whitelists (client_id, ip_address)
			VALUES (?, ?)
		`, clientUUID, ip).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (r *IPWhitelistRepository) FindByClientID(
	ctx context.Context,
	clientUUID string,
) ([]string, error) {
	var ips []string

	err := r.db.WithContext(ctx).
		Raw(`
			SELECT ip_address::text
			FROM client_ip_whitelists
			WHERE client_id = ?
		`, clientUUID).
		Scan(&ips).Error

	return ips, err
}
