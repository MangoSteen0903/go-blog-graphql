package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/MangoSteen0903/go-blog-graphql/ent"
	"github.com/MangoSteen0903/go-blog-graphql/ent/user"
)

// SeeUser is the resolver for the seeUser field.
func (r *queryResolver) SeeUser(ctx context.Context, id int) (*ent.User, error) {
	return r.client.User.Query().Where(user.ID(id)).Only(ctx)
}
