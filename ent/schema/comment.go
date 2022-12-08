package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.String("Context"),
		field.Time("created_at").Default(time.Now()).
			Annotations(
				entgql.OrderField("CREATED_AT"),
			),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("Comments"),
		edge.From("post", Post.Type).Ref("Comments").Annotations(
			entgql.RelayConnection(),
		),
		edge.To("Likes", Like.Type),
	}
}

func (Comment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
