// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/shinnosuke-K/github-dev-insight/ent/repository"
)

// Repository is the model entity for the Repository schema.
type Repository struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// GithubID holds the value of the "github_id" field.
	GithubID string `json:"github_id,omitempty"`
	// Owner holds the value of the "owner" field.
	Owner string `json:"owner,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// TotalPr holds the value of the "total_pr" field.
	TotalPr int64 `json:"total_pr,omitempty"`
	// TotalIssue holds the value of the "total_issue" field.
	TotalIssue int64 `json:"total_issue,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// PushedAt holds the value of the "pushed_at" field.
	PushedAt time.Time `json:"pushed_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RepositoryQuery when eager-loading is set.
	Edges RepositoryEdges `json:"edges"`
}

// RepositoryEdges holds the relations/edges for other nodes in the graph.
type RepositoryEdges struct {
	// PullRequests holds the value of the pull_requests edge.
	PullRequests []*PullRequest `json:"pull_requests,omitempty"`
	// Issues holds the value of the issues edge.
	Issues []*Issue `json:"issues,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// PullRequestsOrErr returns the PullRequests value or an error if the edge
// was not loaded in eager-loading.
func (e RepositoryEdges) PullRequestsOrErr() ([]*PullRequest, error) {
	if e.loadedTypes[0] {
		return e.PullRequests, nil
	}
	return nil, &NotLoadedError{edge: "pull_requests"}
}

// IssuesOrErr returns the Issues value or an error if the edge
// was not loaded in eager-loading.
func (e RepositoryEdges) IssuesOrErr() ([]*Issue, error) {
	if e.loadedTypes[1] {
		return e.Issues, nil
	}
	return nil, &NotLoadedError{edge: "issues"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Repository) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case repository.FieldID, repository.FieldTotalPr, repository.FieldTotalIssue:
			values[i] = new(sql.NullInt64)
		case repository.FieldGithubID, repository.FieldOwner, repository.FieldName, repository.FieldDescription:
			values[i] = new(sql.NullString)
		case repository.FieldCreatedAt, repository.FieldUpdatedAt, repository.FieldPushedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Repository", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Repository fields.
func (r *Repository) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case repository.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int(value.Int64)
		case repository.FieldGithubID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field github_id", values[i])
			} else if value.Valid {
				r.GithubID = value.String
			}
		case repository.FieldOwner:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field owner", values[i])
			} else if value.Valid {
				r.Owner = value.String
			}
		case repository.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				r.Name = value.String
			}
		case repository.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				r.Description = value.String
			}
		case repository.FieldTotalPr:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field total_pr", values[i])
			} else if value.Valid {
				r.TotalPr = value.Int64
			}
		case repository.FieldTotalIssue:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field total_issue", values[i])
			} else if value.Valid {
				r.TotalIssue = value.Int64
			}
		case repository.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				r.CreatedAt = value.Time
			}
		case repository.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				r.UpdatedAt = value.Time
			}
		case repository.FieldPushedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field pushed_at", values[i])
			} else if value.Valid {
				r.PushedAt = value.Time
			}
		}
	}
	return nil
}

// QueryPullRequests queries the "pull_requests" edge of the Repository entity.
func (r *Repository) QueryPullRequests() *PullRequestQuery {
	return (&RepositoryClient{config: r.config}).QueryPullRequests(r)
}

// QueryIssues queries the "issues" edge of the Repository entity.
func (r *Repository) QueryIssues() *IssueQuery {
	return (&RepositoryClient{config: r.config}).QueryIssues(r)
}

// Update returns a builder for updating this Repository.
// Note that you need to call Repository.Unwrap() before calling this method if this Repository
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Repository) Update() *RepositoryUpdateOne {
	return (&RepositoryClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Repository entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Repository) Unwrap() *Repository {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Repository is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Repository) String() string {
	var builder strings.Builder
	builder.WriteString("Repository(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteString(", github_id=")
	builder.WriteString(r.GithubID)
	builder.WriteString(", owner=")
	builder.WriteString(r.Owner)
	builder.WriteString(", name=")
	builder.WriteString(r.Name)
	builder.WriteString(", description=")
	builder.WriteString(r.Description)
	builder.WriteString(", total_pr=")
	builder.WriteString(fmt.Sprintf("%v", r.TotalPr))
	builder.WriteString(", total_issue=")
	builder.WriteString(fmt.Sprintf("%v", r.TotalIssue))
	builder.WriteString(", created_at=")
	builder.WriteString(r.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(r.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", pushed_at=")
	builder.WriteString(r.PushedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Repositories is a parsable slice of Repository.
type Repositories []*Repository

func (r Repositories) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
