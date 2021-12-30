package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Repository holds the schema definition for the Repository entity.
type Repository struct {
	ent.Schema
}

// Fields of the Repository.
func (Repository) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New()).Unique(),
		field.String("github_id").MaxLen(255).NotEmpty(),
		field.String("owner").MaxLen(255).NotEmpty(),
		field.String("name").MaxLen(255).NotEmpty(),
		field.Text("description").Optional(),
		field.Int64("total_pr").Default(0).NonNegative(),
		field.Int64("total_issue").Default(0).NonNegative(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now),
		field.Time("pushed_at").Default(time.Now),
	}
}

// Edges of the Repository.
func (Repository) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("pull_requests", PullRequest.Type),
		edge.To("issues", Issue.Type),
	}
}
