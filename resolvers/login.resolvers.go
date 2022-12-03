package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"os"

	"github.com/MangoSteen0903/go-blog-graphql/ent"
	"github.com/MangoSteen0903/go-blog-graphql/ent/user"
	"github.com/MangoSteen0903/go-blog-graphql/graph/model"
	"github.com/MangoSteen0903/go-blog-graphql/utils"
)

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, username string, password string) (*model.LoginResult, error) {
	userPW := utils.HashingPassword(password)

	user, err := r.client.User.Query().Where(user.Username(username)).Only(ctx)

	if ent.IsNotFound(err) {
		errMsg := fmt.Sprintf("%v", err)
		return &model.LoginResult{
			Ok:    false,
			Error: &errMsg,
		}, nil
	}

	if userPW != user.Password {
		errMsg := "Password does not match. Try again."
		return &model.LoginResult{
			Ok:    false,
			Error: &errMsg,
		}, nil
	}

	newToken := utils.BuildToken(user.ID, username, os.Getenv("JWTKEY"))

	fmt.Println(newToken)
	return &model.LoginResult{
		Ok:    true,
		Token: &newToken,
	}, nil
}
