package middleware

import (
	"context"

	"github.com/MangoSteen0903/go-blog-graphql/ent"
	"github.com/MangoSteen0903/go-blog-graphql/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(client *ent.Client) gin.HandlerFunc {

	return func(c *gin.Context) {
		authToken := c.Request.Header.Get("Token")
		var ctx context.Context
		if authToken == "" {
			ctx = context.WithValue(c.Request.Context(), utils.UserCtxkey, "")
		} else {
			user := utils.GetUser(client, authToken)
			ctx = context.WithValue(c.Request.Context(), utils.UserCtxkey, user)
		}
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
