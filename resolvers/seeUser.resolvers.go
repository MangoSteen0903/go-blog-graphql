package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/MangoSteen0903/go-blog-graphql/ent"
	"github.com/MangoSteen0903/go-blog-graphql/ent/user"
	"github.com/MangoSteen0903/go-blog-graphql/graph/model"
)

// SeeUser is the resolver for the seeUser field.
func (r *queryResolver) SeeUser(ctx context.Context, id int) (*model.UserResult, error) {
	user, err := r.client.User.Query().Where(user.ID(id)).Only(ctx)

	errMsg := fmt.Sprintf("%v", err)
	if ent.IsNotFound(err) {
		return &model.UserResult{
			Ok:    false,
			Error: &errMsg,
		}, nil
	}
	return &model.UserResult{
		Ok:   true,
		User: user,
	}, nil
}
