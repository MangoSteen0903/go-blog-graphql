package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/MangoSteen0903/go-blog-graphql/ent"
	"github.com/MangoSteen0903/go-blog-graphql/graph/model"
	"github.com/MangoSteen0903/go-blog-graphql/utils"
)

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id int, input ent.UpdateUserInput) (*model.Result, error) {
	loggedInUser := utils.ForContext(ctx)

	var errMsg string
	if loggedInUser == nil {
		errMsg = "You need to login to Perform this action. Try again."
		return &model.Result{
			Ok:    false,
			Error: &errMsg,
		}, nil
	}
	if loggedInUser.ID != id {
		errMsg = "Your not authorized to Update this user. Try again."
		return &model.Result{
			Ok:    false,
			Error: &errMsg,
		}, nil
	}

	loggedInUser.Update().SetInput(input).Save(ctx)
	return &model.Result{
		Ok: true,
	}, nil
}
