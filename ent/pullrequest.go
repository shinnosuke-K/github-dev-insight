// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/shinnosuke-K/github-dev-insight/ent/pullrequest"
	"github.com/shinnosuke-K/github-dev-insight/ent/repository"
)

// PullRequest is the model entity for the PullRequest schema.
type PullRequest struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// GithubID holds the value of the "github_id" field.
	GithubID string `json:"github_id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// TotalCommits holds the value of the "total_commits" field.
	TotalCommits int64 `json:"total_commits,omitempty"`
	// GetCommit holds the value of the "get_commit" field.
	// コミット情報を取得したかどうか（0:未取得 1:取得済み）
	GetCommit bool `json:"get_commit,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// ClosedAt holds the value of the "closed_at" field.
	ClosedAt time.Time `json:"closed_at,omitempty"`
	// MergedAt holds the value of the "merged_at" field.
	MergedAt time.Time `json:"merged_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PullRequestQuery when eager-loading is set.
	Edges         PullRequestEdges `json:"edges"`
	repository_id *uuid.UUID
}

// PullRequestEdges holds the relations/edges for other nodes in the graph.
type PullRequestEdges struct {
	// Commits holds the value of the commits edge.
	Commits []*Commits `json:"commits,omitempty"`
	// Repository holds the value of the repository edge.
	Repository *Repository `json:"repository,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// CommitsOrErr returns the Commits value or an error if the edge
// was not loaded in eager-loading.
func (e PullRequestEdges) CommitsOrErr() ([]*Commits, error) {
	if e.loadedTypes[0] {
		return e.Commits, nil
	}
	return nil, &NotLoadedError{edge: "commits"}
}

// RepositoryOrErr returns the Repository value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PullRequestEdges) RepositoryOrErr() (*Repository, error) {
	if e.loadedTypes[1] {
		if e.Repository == nil {
			// The edge repository was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: repository.Label}
		}
		return e.Repository, nil
	}
	return nil, &NotLoadedError{edge: "repository"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PullRequest) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case pullrequest.FieldGetCommit:
			values[i] = new(sql.NullBool)
		case pullrequest.FieldTotalCommits:
			values[i] = new(sql.NullInt64)
		case pullrequest.FieldGithubID, pullrequest.FieldTitle:
			values[i] = new(sql.NullString)
		case pullrequest.FieldCreatedAt, pullrequest.FieldUpdatedAt, pullrequest.FieldClosedAt, pullrequest.FieldMergedAt:
			values[i] = new(sql.NullTime)
		case pullrequest.FieldID:
			values[i] = new(uuid.UUID)
		case pullrequest.ForeignKeys[0]: // repository_id
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type PullRequest", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PullRequest fields.
func (pr *PullRequest) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case pullrequest.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pr.ID = *value
			}
		case pullrequest.FieldGithubID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field github_id", values[i])
			} else if value.Valid {
				pr.GithubID = value.String
			}
		case pullrequest.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				pr.Title = value.String
			}
		case pullrequest.FieldTotalCommits:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field total_commits", values[i])
			} else if value.Valid {
				pr.TotalCommits = value.Int64
			}
		case pullrequest.FieldGetCommit:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field get_commit", values[i])
			} else if value.Valid {
				pr.GetCommit = value.Bool
			}
		case pullrequest.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pr.CreatedAt = value.Time
			}
		case pullrequest.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pr.UpdatedAt = value.Time
			}
		case pullrequest.FieldClosedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field closed_at", values[i])
			} else if value.Valid {
				pr.ClosedAt = value.Time
			}
		case pullrequest.FieldMergedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field merged_at", values[i])
			} else if value.Valid {
				pr.MergedAt = value.Time
			}
		case pullrequest.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field repository_id", values[i])
			} else if value.Valid {
				pr.repository_id = new(uuid.UUID)
				*pr.repository_id = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryCommits queries the "commits" edge of the PullRequest entity.
func (pr *PullRequest) QueryCommits() *CommitsQuery {
	return (&PullRequestClient{config: pr.config}).QueryCommits(pr)
}

// QueryRepository queries the "repository" edge of the PullRequest entity.
func (pr *PullRequest) QueryRepository() *RepositoryQuery {
	return (&PullRequestClient{config: pr.config}).QueryRepository(pr)
}

// Update returns a builder for updating this PullRequest.
// Note that you need to call PullRequest.Unwrap() before calling this method if this PullRequest
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *PullRequest) Update() *PullRequestUpdateOne {
	return (&PullRequestClient{config: pr.config}).UpdateOne(pr)
}

// Unwrap unwraps the PullRequest entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *PullRequest) Unwrap() *PullRequest {
	tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: PullRequest is not a transactional entity")
	}
	pr.config.driver = tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *PullRequest) String() string {
	var builder strings.Builder
	builder.WriteString("PullRequest(")
	builder.WriteString(fmt.Sprintf("id=%v", pr.ID))
	builder.WriteString(", github_id=")
	builder.WriteString(pr.GithubID)
	builder.WriteString(", title=")
	builder.WriteString(pr.Title)
	builder.WriteString(", total_commits=")
	builder.WriteString(fmt.Sprintf("%v", pr.TotalCommits))
	builder.WriteString(", get_commit=")
	builder.WriteString(fmt.Sprintf("%v", pr.GetCommit))
	builder.WriteString(", created_at=")
	builder.WriteString(pr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(pr.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", closed_at=")
	builder.WriteString(pr.ClosedAt.Format(time.ANSIC))
	builder.WriteString(", merged_at=")
	builder.WriteString(pr.MergedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// PullRequests is a parsable slice of PullRequest.
type PullRequests []*PullRequest

func (pr PullRequests) config(cfg config) {
	for _i := range pr {
		pr[_i].config = cfg
	}
}
