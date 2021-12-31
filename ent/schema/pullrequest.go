package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PullRequest holds the schema definition for the PullRequest entity.
type PullRequest struct {
	ent.Schema
}

// Fields of the PullRequest.
func (PullRequest) Fields() []ent.Field {
	return []ent.Field{
		field.String("repository_id").MaxLen(255).NotEmpty(),
		field.String("github_id").MaxLen(255).NotEmpty(),
		field.String("title").MaxLen(255).NotEmpty(),
		field.Int64("total_commits").Default(0).NonNegative(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now),
		field.Time("closed_at").Default(time.Now),
		field.Time("merged_at").Default(time.Now),
	}
}

// Edges of the PullRequest.
func (PullRequest) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("commits", Commits.Type),
		edge.From("repository", Repository.Type).Ref("pull_requests").Unique(),
	}
}
