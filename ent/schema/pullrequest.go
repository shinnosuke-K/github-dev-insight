package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// PullRequest holds the schema definition for the PullRequest entity.
type PullRequest struct {
	ent.Schema
}

// Fields of the PullRequest.
func (PullRequest) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(func() uuid.UUID { return uuid.Must(uuid.NewRandom()) }),
		field.String("github_id").MaxLen(255).NotEmpty(),
		field.String("title").MaxLen(255).NotEmpty(),
		field.Int64("total_commits").Default(0).NonNegative(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now),
		field.Time("closed_at").Default(time.Now).Optional(),
		field.Time("merged_at").Default(time.Now).Optional(),
	}
}

// Edges of the PullRequest.
func (PullRequest) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("commits", Commits.Type).StorageKey(edge.Column("pull_request_id")),
		edge.From("repository", Repository.Type).Ref("pull_requests").Unique(),
	}
}
