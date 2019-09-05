package controllers

import (
	"context"

	"github.com/fitzix/assassin/models"
)

func (r *queryResolver) AllVersions(ctx context.Context, appID *string) ([]*models.Version, error) {
	panic("implement me")
}