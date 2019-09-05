package controllers

import (
	"context"

	"github.com/fitzix/assassin/models"
)

func (r *queryResolver) AllTags(ctx context.Context) ([]*models.Tag, error) {
	panic("implement me")
}