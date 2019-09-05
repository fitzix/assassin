package controllers

import (
	"context"

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


type queryResolver struct{ *Resolver }










