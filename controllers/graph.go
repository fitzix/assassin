package controllers

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/fitzix/assassin/consts"
	"github.com/fitzix/assassin/db"
	"github.com/fitzix/assassin/graph"
	"github.com/fitzix/assassin/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/vektah/gqlparser/gqlerror"
)

// Defining the Graphql handler
func GraphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.GraphQL(graph.NewExecutableSchema(graph.Config{Resolvers: &Resolver{db: db.GetDB(), log: utils.GetLogger()}}))
	return gin.WrapF(h)
}

// Defining the Playground handler
func PlaygroundHandler() gin.HandlerFunc {
	return gin.WrapF(handler.Playground("GraphQL", "/api/query"))
}

func (r *queryResolver) Fail(code consts.StatusCode) error {
	return &gqlerror.Error{
		Message: consts.StatusText(code),
		Extensions: map[string]interface{}{
			"code": code,
		},
	}
}

func (r *queryResolver) Page(query *gorm.DB, size *int, num *int, data interface{}, count interface{}) error {
	if err := query.Model(data).Count(count).Error; err != nil {
		return err
	}
	if err := query.Limit(*size).Offset(*size * (*num - 1)).Find(data).Error; err != nil {
		return err
	}
	return nil
}
