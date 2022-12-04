// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/MangoSteen0903/go-blog-graphql/ent/post"
	"github.com/MangoSteen0903/go-blog-graphql/ent/schema"
	"github.com/MangoSteen0903/go-blog-graphql/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	postFields := schema.Post{}.Fields()
	_ = postFields
	// postDescLikes is the schema descriptor for Likes field.
	postDescLikes := postFields[2].Descriptor()
	// post.DefaultLikes holds the default value on creation for the Likes field.
	post.DefaultLikes = postDescLikes.Default.(int)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUploadImg is the schema descriptor for upload_img field.
	userDescUploadImg := userFields[3].Descriptor()
	// user.DefaultUploadImg holds the default value on creation for the upload_img field.
	user.DefaultUploadImg = userDescUploadImg.Default.(string)
	// userDescIsAdmin is the schema descriptor for is_admin field.
	userDescIsAdmin := userFields[4].Descriptor()
	// user.DefaultIsAdmin holds the default value on creation for the is_admin field.
	user.DefaultIsAdmin = userDescIsAdmin.Default.(bool)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[5].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
}
