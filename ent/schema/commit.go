package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Commit holds the schema definition for the Commit entity.
type Commit struct {
	ent.Schema
}

// Fields of the Commit.
func (Commit) Fields() []ent.Field {
	return []ent.Field{
		field.String("pullrequest_id").MaxLen(255).NotEmpty(),
		field.String("github_id").MaxLen(255).NotEmpty(),
		field.Text("message"),
		field.Time("committed_at").Default(time.Now),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Commit.
func (Commit) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("pull_requests", PullRequest.Type).Ref("commits").Unique(),
	}
}
