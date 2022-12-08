// Code generated by ent, DO NOT EDIT.

package comment

import (
	"time"
)

const (
	// Label holds the string label denoting the comment type in the database.
	Label = "comment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldContext holds the string denoting the context field in the database.
	FieldContext = "context"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgePost holds the string denoting the post edge name in mutations.
	EdgePost = "post"
	// EdgeLikes holds the string denoting the likes edge name in mutations.
	EdgeLikes = "Likes"
	// Table holds the table name of the comment in the database.
	Table = "comments"
	// OwnerTable is the table that holds the owner relation/edge. The primary key declared below.
	OwnerTable = "user_Comments"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// PostTable is the table that holds the post relation/edge. The primary key declared below.
	PostTable = "post_Comments"
	// PostInverseTable is the table name for the Post entity.
	// It exists in this package in order to avoid circular dependency with the "post" package.
	PostInverseTable = "posts"
	// LikesTable is the table that holds the Likes relation/edge. The primary key declared below.
	LikesTable = "comment_Likes"
	// LikesInverseTable is the table name for the Like entity.
	// It exists in this package in order to avoid circular dependency with the "like" package.
	LikesInverseTable = "likes"
)

// Columns holds all SQL columns for comment fields.
var Columns = []string{
	FieldID,
	FieldContext,
	FieldCreatedAt,
}

var (
	// OwnerPrimaryKey and OwnerColumn2 are the table columns denoting the
	// primary key for the owner relation (M2M).
	OwnerPrimaryKey = []string{"user_id", "comment_id"}
	// PostPrimaryKey and PostColumn2 are the table columns denoting the
	// primary key for the post relation (M2M).
	PostPrimaryKey = []string{"post_id", "comment_id"}
	// LikesPrimaryKey and LikesColumn2 are the table columns denoting the
	// primary key for the Likes relation (M2M).
	LikesPrimaryKey = []string{"comment_id", "like_id"}
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt time.Time
)
