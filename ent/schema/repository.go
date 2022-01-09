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
		field.UUID("id", uuid.UUID{}).Default(func() uuid.UUID { return uuid.Must(uuid.NewRandom()) }),
		field.String("github_id").MaxLen(255).NotEmpty(),
		field.String("owner").MaxLen(255).NotEmpty(),
		field.String("name").MaxLen(255).NotEmpty(),
		field.Text("description").Optional(),
		field.Int64("total_pr").Default(0).NonNegative(),
		field.Int64("total_issue").Default(0).NonNegative(),
		field.Bool("get_pull_request").Comment("PR情報を取得したかどうか（0:未取得 1:取得済み）").Default(false),
		field.Bool("get_issue").Comment("Issue情報を取得したかどうか（0:未取得 1:取得済み）").Default(false),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now),
		field.Time("pushed_at").Default(time.Now),
	}
}

// Edges of the Repository.
func (Repository) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("pull_requests", PullRequest.Type).StorageKey(edge.Column("repository_id")),
		edge.To("issues", Issue.Type).StorageKey(edge.Column("repository_id")),
	}
}
