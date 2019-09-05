package controllers

import (
	"context"

	"github.com/fitzix/assassin/models"
)

func (r *mutationResolver) Login(ctx context.Context, input models.Login) (*models.LoginResp, error) {
	panic("implement me")
}