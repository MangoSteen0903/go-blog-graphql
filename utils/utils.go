package utils

import (
	"context"
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/MangoSteen0903/go-blog-graphql/ent"
	"github.com/MangoSteen0903/go-blog-graphql/ent/user"
	"github.com/golang-jwt/jwt/v4"
)

var UserCtxkey = &contextKey{"LoggedInUser"}

type contextKey struct {
	name string
}
type AuthTokenClaims struct {
	UserID   string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func HandleErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s : %v", msg, err)
	}
}

func HashingPassword(str *string) *string {
	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(*str)))
	return &hash
}

func BuildToken(id int, username string, jwtKey string) string {
	tokenStruct := AuthTokenClaims{
		UserID:   fmt.Sprintf("%d", id),
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 5)),
		},
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenStruct)

	newTokenString, err := newToken.SignedString([]byte(jwtKey))

	HandleErr(err, "Can't sign jwt token :")

	return newTokenString

}

func GetUser(client *ent.Client, receiveToken string) *ent.User {

	ctx := context.Background()
	var result *ent.User
	token, err := jwt.ParseWithClaims(receiveToken, &AuthTokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWTKEY")), nil
	})
	HandleErr(err, "JWT token is not valid :")
	if claims, ok := token.Claims.(*AuthTokenClaims); ok && token.Valid {
		userId, err := strconv.Atoi(claims.UserID)
		HandleErr(err, "Can't Convert String to Integer :")
		loggedInUser, queryErr := client.User.Query().Where(user.ID(userId)).Only(ctx)
		HandleErr(queryErr, "Can't find User.")

		result = loggedInUser
	}
	return result
}

func ForContext(ctx context.Context) *ent.User {
	raw, _ := ctx.Value(UserCtxkey).(*ent.User)
	return raw

}
