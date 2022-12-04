package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/MangoSteen0903/go-blog-graphql/ent"
	"github.com/MangoSteen0903/go-blog-graphql/ent/user"
	"github.com/MangoSteen0903/go-blog-graphql/graph/model"
	"github.com/MangoSteen0903/go-blog-graphql/utils"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input ent.CreateUserInput, file *graphql.Upload) (*model.Result, error) {
	newHash := utils.HashingPassword(&input.Password)
	input.Password = *newHash

	if file != nil {
		dirname := "./upload"
		userFileName := fmt.Sprintf("%v/%v-%v-%v.jpg", dirname, input.Username, file.Filename, time.Now())

		newFile, err := os.Create(userFileName)

		if err != nil {
			utils.HandleErr(err, "Error Opening file: ")
		}

		io.Copy(newFile, file.File)
		defer newFile.Close()

		input.UploadImg = &userFileName
	}
	_, err := r.client.User.Create().
		SetInput(input).
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

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id int, input ent.UpdateUserInput, file *graphql.Upload) (*model.Result, error) {
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

	input.Password = utils.HashingPassword(input.Password)
	loggedInUser.Update().SetInput(input).Save(ctx)

	fmt.Println(file)
	return &model.Result{
		Ok: true,
	}, nil
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, username string, password string) (*model.LoginResult, error) {
	userPW := utils.HashingPassword(&password)

	user, err := r.client.User.Query().Where(user.Username(username)).Only(ctx)

	if ent.IsNotFound(err) {
		errMsg := fmt.Sprintf("%v", err)
		return &model.LoginResult{
			Ok:    false,
			Error: &errMsg,
		}, nil
	}

	if *userPW != user.Password {
		errMsg := "Password does not match. Try again."
		return &model.LoginResult{
			Ok:    false,
			Error: &errMsg,
		}, nil
	}

	newToken := utils.BuildToken(user.ID, username, os.Getenv("JWTKEY"))

	return &model.LoginResult{
		Ok:    true,
		Token: &newToken,
	}, nil
}

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

// SeeUsers is the resolver for the seeUsers field.
func (r *queryResolver) SeeUsers(ctx context.Context) ([]*ent.User, error) {
	users, err := r.client.User.Query().All(ctx)
	utils.HandleErr(err, "Can't Query all Users :")
	return users, nil
}
