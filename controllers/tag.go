package controllers

import (
	"context"

	"github.com/fitzix/assassin/models"
)

func (r *queryResolver) AllTags(ctx context.Context) ([]*models.Tag, error) {
	var down []*models.Tag
	if err := r.db.Find(&down).Error; err != nil {
		return nil, r.Fail(3000)
	}
	return down, nil
}