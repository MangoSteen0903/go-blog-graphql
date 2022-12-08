package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/MangoSteen0903/go-blog-graphql/ent"
	"github.com/MangoSteen0903/go-blog-graphql/ent/like"
	"github.com/MangoSteen0903/go-blog-graphql/ent/post"
	"github.com/MangoSteen0903/go-blog-graphql/ent/user"
	"github.com/MangoSteen0903/go-blog-graphql/graph/model"
	"github.com/MangoSteen0903/go-blog-graphql/utils"
)

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input ent.CreatePostInput, hashtags *string) (*model.DefaultResult, error) {
	loggedInUser := utils.ForContext(ctx)
	result := utils.CheckLogin(loggedInUser)

	if result != nil {
		return result, nil
	}

	if !loggedInUser.IsAdmin {
		result := utils.HandleErr("You're not authorized to post on a blog.")
		return &result, nil
	}

	input.OwnerID = &loggedInUser.ID
	if hashtags != nil {
		newHashtags := utils.CreateHashtags(r.client, *hashtags)
		_, err := r.client.Post.Create().
			SetInput(input).
			AddHashtags(newHashtags...).
			Save(ctx)
		if err != nil {
			result := utils.HandleErr("Cannot create Post")
			return &result, nil
		}
	} else {
		_, err := r.client.Post.Create().
			SetInput(input).
			Save(ctx)
		if err != nil {
			result := utils.HandleErr("Cannot create Post")
			return &result, nil
		}
	}

	return &model.DefaultResult{
		Ok: true,
	}, nil
}

// UpdatePost is the resolver for the updatePost field.
func (r *mutationResolver) UpdatePost(ctx context.Context, id int, input ent.UpdatePostInput, hashtags *string) (*model.DefaultResult, error) {
	loggedInUser := utils.ForContext(ctx)

	result := utils.CheckLogin(loggedInUser)

	if result != nil {
		return result, nil
	}

	if !loggedInUser.IsAdmin {
		result := utils.HandleErr("You're not authorized to update on a blog.")
		return &result, nil
	}

	updatedPost, err := r.client.Post.Query().Where(post.ID(id)).Only(ctx)

	if err != nil {
		result := utils.HandleErr("Cannot find Post. Please Try again.")
		return &result, nil
	}

	postOwner, err := updatedPost.QueryOwner().Only(ctx)
	if err != nil {
		result := utils.HandleErr("Cannot find Post's Owner.")
		return &result, nil
	}

	if loggedInUser.ID != postOwner.ID {
		result := utils.HandleErr("You're not authorized to edit this post.")
		return &result, nil
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
			result := utils.HandleErr("Cannot Update Post.")
			return &result, nil
		}
	case newHashtags == nil || hashtags == nil:
		_, err := updatedPost.Update().
			SetInput(input).
			Save(ctx)
		if err != nil {
			result := utils.HandleErr("Cannot Update Post.")
			return &result, nil
		}
	}

	return &model.DefaultResult{
		Ok: true,
	}, nil
}

// DeletePost is the resolver for the deletePost field.
func (r *mutationResolver) DeletePost(ctx context.Context, id int) (*model.DefaultResult, error) {
	loggedInUser := utils.ForContext(ctx)

	result := utils.CheckLogin(loggedInUser)

	if result != nil {
		return result, nil
	}

	post, err := r.client.Post.Query().Where(post.ID(id)).Only(ctx)
	switch {
	case post == nil || err != nil:
		result := utils.HandleErr("Cannot Retrieve post. Please Try again,")
		return &result, nil
	case post != nil:
		owner, _ := post.QueryOwner().Only(ctx)
		if !loggedInUser.IsAdmin || loggedInUser.ID != owner.ID {
			result := utils.HandleErr("You're not authorized to delete post.")
			return &result, nil
		}
		err := r.client.Post.DeleteOneID(id).Exec(ctx)

		if err != nil {
			result := utils.HandleErr("Cannot Delete Post. Please Try again.")
			return &result, nil
		}
	}

	return &model.DefaultResult{
		Ok: true,
	}, nil
}

// ToggleLike is the resolver for the toggleLike field.
func (r *mutationResolver) ToggleLike(ctx context.Context, id int) (*model.DefaultResult, error) {
	loggedInUser := utils.ForContext(ctx)

	result := utils.CheckLogin(loggedInUser)

	if result != nil {
		return result, nil
	}

	retrievedPost, err := r.client.Post.Query().Where(post.ID(id)).Only(ctx)

	if err != nil {
		errResult := utils.HandleErr("Cannot find Post.")
		return &errResult, nil
	}

	_, err = retrievedPost.QueryLikes().QueryOwner().Where(user.ID(loggedInUser.ID)).Only(ctx)

	if err != nil {
		_, err := r.client.Like.Create().AddPostIDs(retrievedPost.ID).AddOwner(loggedInUser).Save(ctx)
		if err != nil {
			errResult := utils.HandleErr("Cannot create Like.")
			return &errResult, nil
		}
	} else {
		_, err := r.client.Like.Delete().Where(
			like.HasOwnerWith(user.ID(loggedInUser.ID)),
			like.HasPostsWith(post.ID(id)),
		).Exec(ctx)

		if err != nil {
			errResult := utils.HandleErr("Cannot Toggle Like.")
			return &errResult, nil
		}
	}

	return &model.DefaultResult{
		Ok: true,
	}, nil
}

// LikeNum is the resolver for the likeNum field.
func (r *postResolver) LikeNum(ctx context.Context, obj *ent.Post) (int, error) {
	likesNum, err := obj.QueryLikes().Count(ctx)
	if err != nil {
		return 0, err
	}
	return likesNum, nil
}

// SeePost is the resolver for the seePost field.
func (r *queryResolver) SeePost(ctx context.Context, id int) (*model.PostResult, error) {
	var errMsg string
	errResult := &model.PostResult{Ok: false}
	post, err := r.client.Post.Query().Where(post.ID(id)).Only(ctx)

	if err != nil {
		errMsg = "Cannot Find Post. Please Try again."
		errResult.Error = &errMsg
		return errResult, nil
	}

	return &model.PostResult{
		Ok:   true,
		Post: post,
	}, nil
}

// SeeUserPost is the resolver for the seeUserPost field.
func (r *queryResolver) SeeUserPost(ctx context.Context, userID int) (*model.PostsResult, error) {
	var errMsg string
	errResult := &model.PostsResult{Ok: false}

	posts, err := r.client.Post.Query().Where(
		post.HasOwnerWith(
			user.ID(userID),
		),
	).All(ctx)

	if err != nil {
		errMsg = "Cannot find User's Post"
		errResult.Error = &errMsg
		return errResult, nil
	}

	return &model.PostsResult{
		Ok:    true,
		Posts: posts,
	}, nil
}
