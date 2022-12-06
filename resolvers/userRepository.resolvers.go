package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/MangoSteen0903/go-blog-graphql/ent"
	"github.com/MangoSteen0903/go-blog-graphql/ent/user"
	"github.com/MangoSteen0903/go-blog-graphql/graph/model"
	"github.com/MangoSteen0903/go-blog-graphql/server/awsLoader"
	"github.com/MangoSteen0903/go-blog-graphql/utils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input ent.CreateUserInput, file *graphql.Upload) (model.Result, error) {
	newHash := utils.HashingPassword(&input.Password)
	input.Password = *newHash

	if file != nil {
		dirname := "./upload"
		userFileName := fmt.Sprintf("%v/%v-%v-%v.jpg", dirname, input.Username, file.Filename, time.Now())

		newFile, err := os.Create(userFileName)

		if err != nil {
			result := utils.HandleErr("Error Opening file")
			return result, nil
		}

		io.Copy(newFile, file.File)

		defer newFile.Close()

		client := awsLoader.LoadAWS()
		_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String("go-blog-bucket"),
			Key:    aws.String(userFileName),
			Body:   newFile,
		})
		if err != nil {
			log.Println("Cannot Upload File.")
		}
		input.UploadImg = &userFileName
	}
	_, err := r.client.User.Create().
		SetInput(input).
		Save(ctx)

	if ent.IsConstraintError(err) {
		result := utils.HandleErr("User name is already Taken. Please Try again.")
		return result, nil
	}

	return &model.DefaultResult{
		Ok:    true,
		Error: nil,
	}, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id int, input ent.UpdateUserInput, file *graphql.Upload) (model.Result, error) {
	loggedInUser := utils.ForContext(ctx)
	if loggedInUser == nil {
		result := utils.HandleErr("You have to login to perform this action.")
		return result, nil
	}
	if loggedInUser.ID != id {
		result := utils.HandleErr("Your not authorized to Update this user. Try again.")
		return result, nil
	}

	input.Password = utils.HashingPassword(input.Password)
	loggedInUser.Update().SetInput(input).Save(ctx)

	return &model.DefaultResult{
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
		result := utils.HandleErr("Password does not match. Try again.")
		newResult := model.LoginResult{Ok: result.Ok, Error: result.Error}
		return &newResult, nil
	}

	newToken, err := utils.BuildToken(user.ID, username, os.Getenv("JWTKEY"))

	if err != nil {
		result := utils.HandleErr("Cannot Create Token. Please Try again.")
		newResult := model.LoginResult{Ok: result.Ok, Error: result.Error}
		return &newResult, nil
	}

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
func (r *queryResolver) SeeUsers(ctx context.Context) (*model.UsersResult, error) {
	users, err := r.client.User.Query().All(ctx)
	if err != nil {
		result := utils.HandleErr("Cannot retrive User. Please Try again.")
		newResult := &model.UsersResult{Ok: result.Ok, Error: result.Error}
		return newResult, nil
	}
	return &model.UsersResult{
		Ok:    true,
		Users: users,
	}, nil
}
