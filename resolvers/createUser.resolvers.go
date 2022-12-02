package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"crypto/sha256"
	"fmt"

	"github.com/MangoSteen0903/go-blog-graphql/ent"
	"github.com/MangoSteen0903/go-blog-graphql/graph/generated"
	"github.com/MangoSteen0903/go-blog-graphql/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input ent.CreateUserInput) (*model.Result, error) {
	newHash := fmt.Sprintf("%x", sha256.Sum256([]byte(input.Password)))
	_, err := r.client.User.Create().
		SetUsername(input.Username).
		SetLocation(*input.Location).
		SetPassword(newHash).
		SetIsAdmin(*input.IsAdmin).
		Save(ctx)

	errMsg := fmt.Sprintf("%v", err)

	if ent.IsConstraintError(err) {
		return &model.Result{
			Ok:    false,
			Error: &errMsg,
		}, nil
	}

	return &model.Result{
		Ok:    true,
		Error: nil,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
