package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/MangoSteen0903/go-blog-graphql/ent"
	"github.com/MangoSteen0903/go-blog-graphql/ent/post"
	"github.com/MangoSteen0903/go-blog-graphql/ent/user"
	"github.com/MangoSteen0903/go-blog-graphql/graph/generated"
	"github.com/MangoSteen0903/go-blog-graphql/graph/model"
	"github.com/MangoSteen0903/go-blog-graphql/utils"
)

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input ent.CreatePostInput, hashtags *string) (*model.Result, error) {
	loggedInUser := utils.ForContext(ctx)

	result := utils.CheckLogin(loggedInUser)

	if result != nil {
		return result, nil
	}

	errMsg := "You're not authorized to post on a blog."
	if !loggedInUser.IsAdmin {
		return &model.Result{
			Ok:    false,
			Error: &errMsg,
		}, nil
	}

	newHashtags := utils.CreateHashtags(r.client, *hashtags)

	input.OwnerID = &loggedInUser.ID
	_, err := r.client.Post.Create().
		SetInput(input).
		AddHashtags(newHashtags...).
		Save(ctx)

	utils.HandleErr(err, "Can't create Post :")

	return &model.Result{
		Ok: true,
	}, nil
}

// SeeUserPost is the resolver for the seeUserPost field.
func (r *queryResolver) SeeUserPost(ctx context.Context, userID int) ([]*ent.Post, error) {
	posts, err := r.client.Post.Query().Where(
		post.HasOwnerWith(
			user.ID(userID),
		),
	).All(ctx)

	utils.HandleErr(err, "Can't find User's Post :")

	fmt.Println(posts)
	return posts, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
