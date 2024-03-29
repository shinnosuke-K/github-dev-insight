// Code generated by ent, DO NOT EDIT.

package pullrequest

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/shinnosuke-K/github-dev-insight/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldLTE(FieldID, id))
}

// GithubID applies equality check predicate on the "github_id" field. It's identical to GithubIDEQ.
func GithubID(v string) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldEQ(FieldGithubID, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldEQ(FieldTitle, v))
}

// TotalCommits applies equality check predicate on the "total_commits" field. It's identical to TotalCommitsEQ.
func TotalCommits(v int64) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldEQ(FieldTotalCommits, v))
}

// GetCommit applies equality check predicate on the "get_commit" field. It's identical to GetCommitEQ.
func GetCommit(v bool) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldEQ(FieldGetCommit, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldEQ(FieldUpdatedAt, v))
}

// ClosedAt applies equality check predicate on the "closed_at" field. It's identical to ClosedAtEQ.
func ClosedAt(v time.Time) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldEQ(FieldClosedAt, v))
}

// MergedAt applies equality check predicate on the "merged_at" field. It's identical to MergedAtEQ.
func MergedAt(v time.Time) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldEQ(FieldMergedAt, v))
}

// GithubIDEQ applies the EQ predicate on the "github_id" field.
func GithubIDEQ(v string) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldEQ(FieldGithubID, v))
}

// GithubIDNEQ applies the NEQ predicate on the "github_id" field.
func GithubIDNEQ(v string) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldNEQ(FieldGithubID, v))
}

// GithubIDIn applies the In predicate on the "github_id" field.
func GithubIDIn(vs ...string) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldIn(FieldGithubID, vs...))
}

// GithubIDNotIn applies the NotIn predicate on the "github_id" field.
func GithubIDNotIn(vs ...string) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldNotIn(FieldGithubID, vs...))
}

// GithubIDGT applies the GT predicate on the "github_id" field.
func GithubIDGT(v string) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldGT(FieldGithubID, v))
}

// GithubIDGTE applies the GTE predicate on the "github_id" field.
func GithubIDGTE(v string) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldGTE(FieldGithubID, v))
}

// GithubIDLT applies the LT predicate on the "github_id" field.
func GithubIDLT(v string) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldLT(FieldGithubID, v))
}

// GithubIDLTE applies the LTE predicate on the "github_id" field.
func GithubIDLTE(v string) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldLTE(FieldGithubID, v))
}

// GithubIDContains applies the Contains predicate on the "github_id" field.
func GithubIDContains(v string) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldContains(FieldGithubID, v))
}

// GithubIDHasPrefix applies the HasPrefix predicate on the "github_id" field.
func GithubIDHasPrefix(v string) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldHasPrefix(FieldGithubID, v))
}

// GithubIDHasSuffix applies the HasSuffix predicate on the "github_id" field.
func GithubIDHasSuffix(v string) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldHasSuffix(FieldGithubID, v))
}

// GithubIDEqualFold applies the EqualFold predicate on the "github_id" field.
func GithubIDEqualFold(v string) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldEqualFold(FieldGithubID, v))
}

// GithubIDContainsFold applies the ContainsFold predicate on the "github_id" field.
func GithubIDContainsFold(v string) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldContainsFold(FieldGithubID, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldContainsFold(FieldTitle, v))
}

// TotalCommitsEQ applies the EQ predicate on the "total_commits" field.
func TotalCommitsEQ(v int64) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldEQ(FieldTotalCommits, v))
}

// TotalCommitsNEQ applies the NEQ predicate on the "total_commits" field.
func TotalCommitsNEQ(v int64) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldNEQ(FieldTotalCommits, v))
}

// TotalCommitsIn applies the In predicate on the "total_commits" field.
func TotalCommitsIn(vs ...int64) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldIn(FieldTotalCommits, vs...))
}

// TotalCommitsNotIn applies the NotIn predicate on the "total_commits" field.
func TotalCommitsNotIn(vs ...int64) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldNotIn(FieldTotalCommits, vs...))
}

// TotalCommitsGT applies the GT predicate on the "total_commits" field.
func TotalCommitsGT(v int64) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldGT(FieldTotalCommits, v))
}

// TotalCommitsGTE applies the GTE predicate on the "total_commits" field.
func TotalCommitsGTE(v int64) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldGTE(FieldTotalCommits, v))
}

// TotalCommitsLT applies the LT predicate on the "total_commits" field.
func TotalCommitsLT(v int64) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldLT(FieldTotalCommits, v))
}

// TotalCommitsLTE applies the LTE predicate on the "total_commits" field.
func TotalCommitsLTE(v int64) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldLTE(FieldTotalCommits, v))
}

// GetCommitEQ applies the EQ predicate on the "get_commit" field.
func GetCommitEQ(v bool) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldEQ(FieldGetCommit, v))
}

// GetCommitNEQ applies the NEQ predicate on the "get_commit" field.
func GetCommitNEQ(v bool) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldNEQ(FieldGetCommit, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldLTE(FieldUpdatedAt, v))
}

// ClosedAtEQ applies the EQ predicate on the "closed_at" field.
func ClosedAtEQ(v time.Time) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldEQ(FieldClosedAt, v))
}

// ClosedAtNEQ applies the NEQ predicate on the "closed_at" field.
func ClosedAtNEQ(v time.Time) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldNEQ(FieldClosedAt, v))
}

// ClosedAtIn applies the In predicate on the "closed_at" field.
func ClosedAtIn(vs ...time.Time) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldIn(FieldClosedAt, vs...))
}

// ClosedAtNotIn applies the NotIn predicate on the "closed_at" field.
func ClosedAtNotIn(vs ...time.Time) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldNotIn(FieldClosedAt, vs...))
}

// ClosedAtGT applies the GT predicate on the "closed_at" field.
func ClosedAtGT(v time.Time) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldGT(FieldClosedAt, v))
}

// ClosedAtGTE applies the GTE predicate on the "closed_at" field.
func ClosedAtGTE(v time.Time) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldGTE(FieldClosedAt, v))
}

// ClosedAtLT applies the LT predicate on the "closed_at" field.
func ClosedAtLT(v time.Time) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldLT(FieldClosedAt, v))
}

// ClosedAtLTE applies the LTE predicate on the "closed_at" field.
func ClosedAtLTE(v time.Time) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldLTE(FieldClosedAt, v))
}

// ClosedAtIsNil applies the IsNil predicate on the "closed_at" field.
func ClosedAtIsNil() predicate.PullRequest {
	return predicate.PullRequest(sql.FieldIsNull(FieldClosedAt))
}

// ClosedAtNotNil applies the NotNil predicate on the "closed_at" field.
func ClosedAtNotNil() predicate.PullRequest {
	return predicate.PullRequest(sql.FieldNotNull(FieldClosedAt))
}

// MergedAtEQ applies the EQ predicate on the "merged_at" field.
func MergedAtEQ(v time.Time) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldEQ(FieldMergedAt, v))
}

// MergedAtNEQ applies the NEQ predicate on the "merged_at" field.
func MergedAtNEQ(v time.Time) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldNEQ(FieldMergedAt, v))
}

// MergedAtIn applies the In predicate on the "merged_at" field.
func MergedAtIn(vs ...time.Time) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldIn(FieldMergedAt, vs...))
}

// MergedAtNotIn applies the NotIn predicate on the "merged_at" field.
func MergedAtNotIn(vs ...time.Time) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldNotIn(FieldMergedAt, vs...))
}

// MergedAtGT applies the GT predicate on the "merged_at" field.
func MergedAtGT(v time.Time) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldGT(FieldMergedAt, v))
}

// MergedAtGTE applies the GTE predicate on the "merged_at" field.
func MergedAtGTE(v time.Time) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldGTE(FieldMergedAt, v))
}

// MergedAtLT applies the LT predicate on the "merged_at" field.
func MergedAtLT(v time.Time) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldLT(FieldMergedAt, v))
}

// MergedAtLTE applies the LTE predicate on the "merged_at" field.
func MergedAtLTE(v time.Time) predicate.PullRequest {
	return predicate.PullRequest(sql.FieldLTE(FieldMergedAt, v))
}

// MergedAtIsNil applies the IsNil predicate on the "merged_at" field.
func MergedAtIsNil() predicate.PullRequest {
	return predicate.PullRequest(sql.FieldIsNull(FieldMergedAt))
}

// MergedAtNotNil applies the NotNil predicate on the "merged_at" field.
func MergedAtNotNil() predicate.PullRequest {
	return predicate.PullRequest(sql.FieldNotNull(FieldMergedAt))
}

// HasCommits applies the HasEdge predicate on the "commits" edge.
func HasCommits() predicate.PullRequest {
	return predicate.PullRequest(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CommitsTable, CommitsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCommitsWith applies the HasEdge predicate on the "commits" edge with a given conditions (other predicates).
func HasCommitsWith(preds ...predicate.Commits) predicate.PullRequest {
	return predicate.PullRequest(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CommitsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CommitsTable, CommitsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRepository applies the HasEdge predicate on the "repository" edge.
func HasRepository() predicate.PullRequest {
	return predicate.PullRequest(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RepositoryTable, RepositoryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRepositoryWith applies the HasEdge predicate on the "repository" edge with a given conditions (other predicates).
func HasRepositoryWith(preds ...predicate.Repository) predicate.PullRequest {
	return predicate.PullRequest(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RepositoryInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RepositoryTable, RepositoryColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PullRequest) predicate.PullRequest {
	return predicate.PullRequest(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PullRequest) predicate.PullRequest {
	return predicate.PullRequest(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PullRequest) predicate.PullRequest {
	return predicate.PullRequest(func(s *sql.Selector) {
		p(s.Not())
	})
}
