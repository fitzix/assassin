package controllers

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/fitzix/assassin/graph"
	"github.com/gin-gonic/gin"
)

// Defining the Graphql handler
func GraphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.GraphQL(graph.NewExecutableSchema(graph.Config{Resolvers: &Resolver{}}))
	return gin.WrapF(h)
}

// Defining the Playground handler
func PlaygroundHandler() gin.HandlerFunc {
	return gin.WrapF(handler.Playground("GraphQL", "/api/query"))
}
