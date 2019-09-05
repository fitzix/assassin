package controllers

import (
	"context"
	"fmt"

	"github.com/fitzix/assassin/models"
)

func (r *queryResolver) AllApps(ctx context.Context, key *string, size *int, num *int, order *models.AppOrderBy) (*models.AppList, error) {
	var down models.AppList
	d := r.db
	if key != nil {
		d = d.Where("name LIKE ?", fmt.Sprintf("%%%s%%", *key))
	}

	switch *order {
	case models.AppOrderByHot:
		d = d.Joins("LEFT JOIN app_hot ON app.id = app_hot.app_id").Order("app_hot.hot DESC")
	case models.AppOrderByUpdateDesc:
		d = d.Order("updated_at DESC")
	}

	if err := r.Page(d, size, num, &down.Apps, &down.Total); err != nil {
		return nil, r.Fail(3000)
	}
	return &down, nil
}

func (r *queryResolver) App(ctx context.Context, id string) (*models.App, error) {
	var down models.App
	if err := r.db.Find(&down, "id = ?", id).Error; err != nil {
		r.log.Errorf("db err: %s", err)
		return nil, r.Fail(3000)
	}
	return &down, nil
}
