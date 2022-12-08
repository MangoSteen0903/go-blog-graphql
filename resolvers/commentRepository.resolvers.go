package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/MangoSteen0903/go-blog-graphql/ent"
	"github.com/MangoSteen0903/go-blog-graphql/ent/comment"
	"github.com/MangoSteen0903/go-blog-graphql/ent/like"
	"github.com/MangoSteen0903/go-blog-graphql/ent/post"
	"github.com/MangoSteen0903/go-blog-graphql/ent/user"
	"github.com/MangoSteen0903/go-blog-graphql/graph/generated"
	"github.com/MangoSteen0903/go-blog-graphql/graph/model"
	"github.com/MangoSteen0903/go-blog-graphql/utils"
)

// LikeNum is the resolver for the likeNum field.
func (r *commentResolver) LikeNum(ctx context.Context, obj *ent.Comment) (int, error) {
	count, err := obj.QueryLikes().Count(ctx)
	if err != nil {
		return 0, nil
	}
	return count, nil
}

// PostComment is the resolver for the postComment field.
func (r *mutationResolver) PostComment(ctx context.Context, postID int, input ent.CreateCommentInput) (*model.DefaultResult, error) {
	loggedInUser := utils.ForContext(ctx)

	result := utils.CheckLogin(loggedInUser)

	if result != nil {
		return result, nil
	}

	_, err := r.client.Post.Query().Where(post.ID(postID)).Only(ctx)

	if err != nil {
		errResult := utils.HandleErr("Can't find Post.")
		return &errResult, nil
	}

	_, err = r.client.Comment.Create().
		SetInput(input).
		AddPostIDs(postID).
		AddOwnerIDs(loggedInUser.ID).
		Save(ctx)

	if err != nil {
		errResult := utils.HandleErr("Can't create Comment")
		return &errResult, nil
	}
	return &model.DefaultResult{
		Ok: true,
	}, nil
}

// EditComment is the resolver for the editComment field.
func (r *mutationResolver) EditComment(ctx context.Context, id int, input ent.UpdateCommentInput) (*model.DefaultResult, error) {
	loggedInUser := utils.ForContext(ctx)

	result := utils.CheckLogin(loggedInUser)

	if result != nil {
		return result, nil
	}

	comment, err := r.client.Comment.Query().Where(comment.ID(id)).Only(ctx)

	if err != nil {
		errResult := utils.HandleErr("Can't find Comment.")
		return &errResult, nil
	}

	owner, _ := comment.QueryOwner().Unique(true).Only(ctx)

	if owner.ID != loggedInUser.ID {
		errResult := utils.HandleErr("You're not authorized to edit this comment")
		return &errResult, nil
	}

	err = comment.Update().SetInput(input).Exec(ctx)

	if err != nil {
		errResult := utils.HandleErr("Can't update Comment.")
		return &errResult, nil
	}
	return &model.DefaultResult{
		Ok: true,
	}, nil
}

// DeleteComment is the resolver for the deleteComment field.
func (r *mutationResolver) DeleteComment(ctx context.Context, id int) (*model.DefaultResult, error) {
	loggedInUser := utils.ForContext(ctx)

	result := utils.CheckLogin(loggedInUser)

	if result != nil {
		return result, nil
	}

	comment, err := r.client.Comment.Query().Where(comment.ID(id)).Only(ctx)

	switch {
	case comment == nil || err != nil:
		errResult := utils.HandleErr("Can't find Comment.")
		return &errResult, nil
	case comment != nil:
		owner, _ := comment.QueryOwner().Only(ctx)

		if owner.ID != loggedInUser.ID {
			errResult := utils.HandleErr("You're not authorized to Delete this post.")
			return &errResult, nil
		}
		err = r.client.Comment.DeleteOneID(id).Exec(ctx)
		if err != nil {
			errResult := utils.HandleErr("Can't delete Comment.")
			return &errResult, nil
		}
	}

	return &model.DefaultResult{
		Ok: true,
	}, nil
}

// ToggleCommentLike is the resolver for the toggleCommentLike field.
func (r *mutationResolver) ToggleCommentLike(ctx context.Context, id int) (*model.DefaultResult, error) {
	loggedInUser := utils.ForContext(ctx)

	result := utils.CheckLogin(loggedInUser)

	if result != nil {
		return result, nil
	}

	retrievedComment, err := r.client.Comment.Query().Where(comment.ID(id)).Only(ctx)

	if err != nil {
		errResult := utils.HandleErr("Can't find comment. Please Try again.")
		return &errResult, nil
	}

	_, err = retrievedComment.QueryLikes().QueryOwner().Where(user.ID(loggedInUser.ID)).Only(ctx)

	if err != nil {
		_, err = r.client.Like.Create().AddCommentIDs(id).AddOwnerIDs(loggedInUser.ID).Save(ctx)
		if err != nil {
			errResult := utils.HandleErr("Can't create Like. Please Try again.")
			return &errResult, nil
		}
	} else {
		_, err = r.client.Like.Delete().Where(
			like.HasCommentsWith(comment.ID(id)),
			like.HasOwnerWith(user.ID(loggedInUser.ID)),
		).Exec(ctx)

		if err != nil {
			errResult := utils.HandleErr("Can't delete Comment. Please try again.")
			return &errResult, nil
		}
	}
	return &model.DefaultResult{
		Ok: true,
	}, nil
}

// SeePostComment is the resolver for the seePostComment field.
func (r *queryResolver) SeePostComment(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.CommentOrder, postID int) (*ent.CommentConnection, error) {
	comments, err := r.client.Post.Query().Where(post.ID(postID)).QueryComments().Paginate(ctx, after, first, before, last,
		ent.WithCommentOrder(orderBy),
	)
	if err != nil {
		return nil, err
	}

	fmt.Println(comments)

	return comments, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
