package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Unique(),
		field.String("password"),
		field.String("location").Optional(),
		field.String("upload_img").Optional().Default(""),
		field.Bool("is_admin").Default(false),
		field.Time("created_at").Default(time.Now).
			Annotations(
				entgql.OrderField("CREATED_AT"),
			),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("Posts", Post.Type).Annotations(
			entgql.RelayConnection(),
		),
		edge.To("Likes", Like.Type),
		edge.To("Comments", Comment.Type).Annotations(
			entgql.RelayConnection(),
		),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
