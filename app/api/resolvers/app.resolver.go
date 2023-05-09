package resolvers

import (
	"geoloc/app/api/graphql/generated/graph"
	"geoloc/app/services"
)

type Resolver struct {
	services *services.Service
}

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

func NewResolver(s *services.Service) *Resolver {
	return &Resolver{s}
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }
