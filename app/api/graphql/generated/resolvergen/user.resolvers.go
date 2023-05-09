package resolvergen

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"geoloc/app/api/graphql/generated/graph"
	"geoloc/app/models"
)

// SuperAdminCreate is the resolver for the superAdminCreate field.
func (r *mutationResolver) SuperAdminCreate(ctx context.Context, input graph.UpdateUser) (*models.User, error) {
	panic(fmt.Errorf("not implemented: SuperAdminCreate - superAdminCreate"))
}

// UserCreate is the resolver for the userCreate field.
func (r *mutationResolver) UserCreate(ctx context.Context, input graph.UpdateUser) (*models.User, error) {
	panic(fmt.Errorf("not implemented: UserCreate - userCreate"))
}

// UserUpdate is the resolver for the userUpdate field.
func (r *mutationResolver) UserUpdate(ctx context.Context, id int64, input graph.UpdateUser) (*models.User, error) {
	panic(fmt.Errorf("not implemented: UserUpdate - userUpdate"))
}

// ResendEmailVerification is the resolver for the resendEmailVerification field.
func (r *mutationResolver) ResendEmailVerification(ctx context.Context, email string) (bool, error) {
	panic(fmt.Errorf("not implemented: ResendEmailVerification - resendEmailVerification"))
}

// UserArchive is the resolver for the userArchive field.
func (r *mutationResolver) UserArchive(ctx context.Context, id int64) (*models.User, error) {
	panic(fmt.Errorf("not implemented: UserArchive - userArchive"))
}

// UserUnarchive is the resolver for the userUnarchive field.
func (r *mutationResolver) UserUnarchive(ctx context.Context, id int64) (*models.User, error) {
	panic(fmt.Errorf("not implemented: UserUnarchive - userUnarchive"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, search graph.SearchFilter) (*graph.UserResult, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id *int64, email *string) (*models.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*models.User, error) {
	panic(fmt.Errorf("not implemented: Me - me"))
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
