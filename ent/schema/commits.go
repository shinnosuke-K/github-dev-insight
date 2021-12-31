package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Commit holds the schema definition for the Commits entity.
type Commits struct {
	ent.Schema
}

// Fields of the Commit.
func (Commits) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(func() uuid.UUID { return uuid.Must(uuid.NewRandom()) }),
		field.String("github_id").MaxLen(255).NotEmpty(),
		field.Text("message"),
		field.Time("committed_at").Default(time.Now),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Commit.
func (Commits) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("pull_request", PullRequest.Type).Ref("commits").Unique(),
	}
}
