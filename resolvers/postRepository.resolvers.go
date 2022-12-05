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
	var errMsg string
	errResult := model.Result{Ok: false}

	loggedInUser := utils.ForContext(ctx)
	result := utils.CheckLogin(loggedInUser)

	if result != nil {
		return result, nil
	}

	if !loggedInUser.IsAdmin {
		errMsg = "You're not authorized to post on a blog."
		errResult.Error = &errMsg
		return &errResult, nil
	}

	input.OwnerID = &loggedInUser.ID
	if hashtags != nil {
		newHashtags := utils.CreateHashtags(r.client, *hashtags)
		_, err := r.client.Post.Create().
			SetInput(input).
			AddHashtags(newHashtags...).
			Save(ctx)
		if err != nil {
			errMsg = "Can't Create Post."
			errResult.Error = &errMsg
			return &errResult, nil
		}
	} else {
		_, err := r.client.Post.Create().
			SetInput(input).
			Save(ctx)
		if err != nil {
			errMsg = "Can't Create Post."
			errResult.Error = &errMsg
			return &errResult, nil
		}
	}

	return &model.Result{
		Ok: true,
	}, nil
}

// UpdatePost is the resolver for the updatePost field.
func (r *mutationResolver) UpdatePost(ctx context.Context, id int, input ent.UpdatePostInput, hashtags *string) (*model.Result, error) {
	var errMsg string
	errResult := model.Result{Ok: false}

	loggedInUser := utils.ForContext(ctx)

	result := utils.CheckLogin(loggedInUser)

	if result != nil {
		return result, nil
	}

	if !loggedInUser.IsAdmin {
		errMsg = "You're not authorized to update on a blog."
		errResult.Error = &errMsg
		return &errResult, nil
	}

	updatedPost, err := r.client.Post.Query().Where(post.ID(id)).Only(ctx)

	if err != nil {
		errMsg = "Cannot find Post. Please Try again."
		errResult.Error = &errMsg
		return &errResult, nil
	}

	postOwner, err := updatedPost.QueryOwner().Only(ctx)
	if err != nil {
		errMsg = "Cannot find Post's Owner."
		errResult.Error = &errMsg
		return &errResult, nil
	}

	if loggedInUser.ID != postOwner.ID {
		errMsg = "Your not authorized to Edit This Post."
		errResult.Error = &errMsg
		return &errResult, nil
	}

	var newHashtags []*ent.Hashtag
	if hashtags != nil {
		newHashtags = utils.CreateHashtags(r.client, *hashtags)
	}

	switch {
	case newHashtags != nil:
		_, err := updatedPost.Update().
			SetInput(input).
			AddHashtags(newHashtags...).
			Save(ctx)
		if err != nil {
			errMsg = "Can't Update Post"
			errResult.Error = &errMsg
			return &errResult, nil
		}
	case newHashtags == nil || hashtags == nil:
		_, err := updatedPost.Update().
			SetInput(input).
			Save(ctx)
		if err != nil {
			errMsg = "Can't Update Post"
			errResult.Error = &errMsg
			return &errResult, nil
		}
	}

	return &model.Result{
		Ok: true,
	}, nil
}

// DeletePost is the resolver for the deletePost field.
func (r *mutationResolver) DeletePost(ctx context.Context, id int) (*model.Result, error) {

	var errMsg string
	errResult := &model.Result{Ok: false}

	loggedInUser := utils.ForContext(ctx)

	result := utils.CheckLogin(loggedInUser)

	if result != nil {
		return result, nil
	}

	post, err := r.client.Post.Query().Where(post.ID(id)).Only(ctx)
	switch {
	case post == nil || err != nil:
		errMsg = "Cannot Retrieve post. Please Try again,"
		errResult.Error = &errMsg
		return errResult, nil
	case post != nil:
		owner, _ := post.QueryOwner().Only(ctx)
		if !loggedInUser.IsAdmin || loggedInUser.ID != owner.ID {
			errMsg = "You're not authorized to delete post."
			errResult.Error = &errMsg
			return errResult, nil
		}
		err := r.client.Post.DeleteOneID(id).Exec(ctx)

		if err != nil {
			errMsg = "Cannot Delete Post. Please Try again."
			errResult.Error = &errMsg
			return errResult, nil
		}
	}

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

	utils.HandleErr(err, "Can't find User's Post.")

	fmt.Println(posts)
	return posts, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
