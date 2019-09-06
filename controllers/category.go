package controllers

import (
	"context"

	"github.com/fitzix/assassin/models"
)

func (r *queryResolver) AllCategories(ctx context.Context) ([]*models.Category, error) {
	var down []*models.Category
	if err := r.db.Find(&down).Error; err != nil {
		return nil, r.Fail(3000)
	}
	return down, nil
}
