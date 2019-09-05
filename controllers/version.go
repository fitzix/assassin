package controllers

import (
	"context"

	"github.com/fitzix/assassin/models"
)

func (r *queryResolver) AllVersions(ctx context.Context, appID string) ([]*models.AppVersion, error) {
	var down []*models.AppVersion
	if err := r.db.Order("created_at DESC").Find(&down, "app_id = ?", appID).Error; err != nil {
		return nil, r.Fail(3000)
	}
	return down, nil
}