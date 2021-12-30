package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Issue holds the schema definition for the Issue entity.
type Issue struct {
	ent.Schema
}

// Fields of the Issue.
func (Issue) Fields() []ent.Field {
	return []ent.Field{
		field.String("repository_id").MaxLen(255).NotEmpty(),
		field.String("github_id").MaxLen(255).NotEmpty(),
		field.String("title").MaxLen(255).NotEmpty(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now),
		field.Time("last_edited_at").Default(time.Now),
		field.Time("closed_at").Default(time.Now),
	}
}

// Edges of the Issue.
func (Issue) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("repository", Repository.Type).Ref("issues").Unique(),
	}
}
