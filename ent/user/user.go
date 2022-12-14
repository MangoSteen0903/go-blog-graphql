// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldLocation holds the string denoting the location field in the database.
	FieldLocation = "location"
	// FieldUploadImg holds the string denoting the upload_img field in the database.
	FieldUploadImg = "upload_img"
	// FieldIsAdmin holds the string denoting the is_admin field in the database.
	FieldIsAdmin = "is_admin"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgePosts holds the string denoting the posts edge name in mutations.
	EdgePosts = "Posts"
	// EdgeLikes holds the string denoting the likes edge name in mutations.
	EdgeLikes = "Likes"
	// EdgeComments holds the string denoting the comments edge name in mutations.
	EdgeComments = "Comments"
	// Table holds the table name of the user in the database.
	Table = "users"
	// PostsTable is the table that holds the Posts relation/edge.
	PostsTable = "posts"
	// PostsInverseTable is the table name for the Post entity.
	// It exists in this package in order to avoid circular dependency with the "post" package.
	PostsInverseTable = "posts"
	// PostsColumn is the table column denoting the Posts relation/edge.
	PostsColumn = "user_posts"
	// LikesTable is the table that holds the Likes relation/edge. The primary key declared below.
	LikesTable = "user_Likes"
	// LikesInverseTable is the table name for the Like entity.
	// It exists in this package in order to avoid circular dependency with the "like" package.
	LikesInverseTable = "likes"
	// CommentsTable is the table that holds the Comments relation/edge. The primary key declared below.
	CommentsTable = "user_Comments"
	// CommentsInverseTable is the table name for the Comment entity.
	// It exists in this package in order to avoid circular dependency with the "comment" package.
	CommentsInverseTable = "comments"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUsername,
	FieldPassword,
	FieldLocation,
	FieldUploadImg,
	FieldIsAdmin,
	FieldCreatedAt,
}

var (
	// LikesPrimaryKey and LikesColumn2 are the table columns denoting the
	// primary key for the Likes relation (M2M).
	LikesPrimaryKey = []string{"user_id", "like_id"}
	// CommentsPrimaryKey and CommentsColumn2 are the table columns denoting the
	// primary key for the Comments relation (M2M).
	CommentsPrimaryKey = []string{"user_id", "comment_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultUploadImg holds the default value on creation for the "upload_img" field.
	DefaultUploadImg string
	// DefaultIsAdmin holds the default value on creation for the "is_admin" field.
	DefaultIsAdmin bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)
