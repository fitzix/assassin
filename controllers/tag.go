package controllers

import (
	"context"

	"github.com/fitzix/assassin/db"
	"github.com/fitzix/assassin/models"
)

func (r *queryResolver) AllTags(ctx context.Context) ([]*models.Tag, error) {
	var down []*models.Tag
	d := db.GetDB()
	if err := d.Find(&down).Error; err != nil {
		return nil, r.Fail(3000)
	}
	return down, nil
}