package controllers

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/fitzix/assassin/graph"
	"github.com/fitzix/assassin/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() graph.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() graph.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, input models.NewTodo) (*models.Todo, error) {
	todo := &models.Todo{
		Text:   input.Text,
		ID:     fmt.Sprintf("T%d", rand.Int()),
	}
	return todo, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]*models.Todo, error) {
	ret := []*models.Todo{{
		ID:   "232",
		Text: "22323",
		Done: false,
		User: &models.User{
			ID:   "321",
			Name: "wqeeq",
		},
	}}
	return ret, nil
}
