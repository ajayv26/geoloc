package handlers

import (
	"geoloc/app/api/graphql/generated/graph"
	"geoloc/app/api/resolvers"
	"geoloc/app/services"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

type GraphQLHandler struct {
	services *services.Service
}

// var _ GraphQLHandlerInterface = &GraphQLHandler{}

type GraphQLHandlerInterface interface {
	PlayGround() http.HandlerFunc
	Query() *handler.Server
}

func NewGraphQLHandler(s *services.Service) *GraphQLHandler {
	return &GraphQLHandler{s}
}

func (h *GraphQLHandler) PlayGround() http.HandlerFunc {
	return playground.Handler("GraphQL playground", "/api/gql/query")
}

func (h *GraphQLHandler) Query() *handler.Server {
	resolvers := resolvers.NewResolver(h.services)
	return handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolvers}))
}
