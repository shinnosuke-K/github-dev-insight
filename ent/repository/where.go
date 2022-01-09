// Code generated by entc, DO NOT EDIT.

package repository

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/shinnosuke-K/github-dev-insight/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// GithubID applies equality check predicate on the "github_id" field. It's identical to GithubIDEQ.
func GithubID(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGithubID), v))
	})
}

// Owner applies equality check predicate on the "owner" field. It's identical to OwnerEQ.
func Owner(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOwner), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// TotalPr applies equality check predicate on the "total_pr" field. It's identical to TotalPrEQ.
func TotalPr(v int64) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTotalPr), v))
	})
}

// TotalIssue applies equality check predicate on the "total_issue" field. It's identical to TotalIssueEQ.
func TotalIssue(v int64) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTotalIssue), v))
	})
}

// GetPullRequest applies equality check predicate on the "get_pull_request" field. It's identical to GetPullRequestEQ.
func GetPullRequest(v bool) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGetPullRequest), v))
	})
}

// GetIssue applies equality check predicate on the "get_issue" field. It's identical to GetIssueEQ.
func GetIssue(v bool) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGetIssue), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// PushedAt applies equality check predicate on the "pushed_at" field. It's identical to PushedAtEQ.
func PushedAt(v time.Time) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPushedAt), v))
	})
}

// GithubIDEQ applies the EQ predicate on the "github_id" field.
func GithubIDEQ(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGithubID), v))
	})
}

// GithubIDNEQ applies the NEQ predicate on the "github_id" field.
func GithubIDNEQ(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGithubID), v))
	})
}

// GithubIDIn applies the In predicate on the "github_id" field.
func GithubIDIn(vs ...string) predicate.Repository {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Repository(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldGithubID), v...))
	})
}

// GithubIDNotIn applies the NotIn predicate on the "github_id" field.
func GithubIDNotIn(vs ...string) predicate.Repository {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Repository(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldGithubID), v...))
	})
}

// GithubIDGT applies the GT predicate on the "github_id" field.
func GithubIDGT(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGithubID), v))
	})
}

// GithubIDGTE applies the GTE predicate on the "github_id" field.
func GithubIDGTE(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGithubID), v))
	})
}

// GithubIDLT applies the LT predicate on the "github_id" field.
func GithubIDLT(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGithubID), v))
	})
}

// GithubIDLTE applies the LTE predicate on the "github_id" field.
func GithubIDLTE(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGithubID), v))
	})
}

// GithubIDContains applies the Contains predicate on the "github_id" field.
func GithubIDContains(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldGithubID), v))
	})
}

// GithubIDHasPrefix applies the HasPrefix predicate on the "github_id" field.
func GithubIDHasPrefix(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldGithubID), v))
	})
}

// GithubIDHasSuffix applies the HasSuffix predicate on the "github_id" field.
func GithubIDHasSuffix(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldGithubID), v))
	})
}

// GithubIDEqualFold applies the EqualFold predicate on the "github_id" field.
func GithubIDEqualFold(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldGithubID), v))
	})
}

// GithubIDContainsFold applies the ContainsFold predicate on the "github_id" field.
func GithubIDContainsFold(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldGithubID), v))
	})
}

// OwnerEQ applies the EQ predicate on the "owner" field.
func OwnerEQ(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOwner), v))
	})
}

// OwnerNEQ applies the NEQ predicate on the "owner" field.
func OwnerNEQ(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOwner), v))
	})
}

// OwnerIn applies the In predicate on the "owner" field.
func OwnerIn(vs ...string) predicate.Repository {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Repository(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOwner), v...))
	})
}

// OwnerNotIn applies the NotIn predicate on the "owner" field.
func OwnerNotIn(vs ...string) predicate.Repository {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Repository(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOwner), v...))
	})
}

// OwnerGT applies the GT predicate on the "owner" field.
func OwnerGT(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOwner), v))
	})
}

// OwnerGTE applies the GTE predicate on the "owner" field.
func OwnerGTE(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOwner), v))
	})
}

// OwnerLT applies the LT predicate on the "owner" field.
func OwnerLT(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOwner), v))
	})
}

// OwnerLTE applies the LTE predicate on the "owner" field.
func OwnerLTE(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOwner), v))
	})
}

// OwnerContains applies the Contains predicate on the "owner" field.
func OwnerContains(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOwner), v))
	})
}

// OwnerHasPrefix applies the HasPrefix predicate on the "owner" field.
func OwnerHasPrefix(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOwner), v))
	})
}

// OwnerHasSuffix applies the HasSuffix predicate on the "owner" field.
func OwnerHasSuffix(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOwner), v))
	})
}

// OwnerEqualFold applies the EqualFold predicate on the "owner" field.
func OwnerEqualFold(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOwner), v))
	})
}

// OwnerContainsFold applies the ContainsFold predicate on the "owner" field.
func OwnerContainsFold(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOwner), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Repository {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Repository(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Repository {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Repository(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescription), v))
	})
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Repository {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Repository(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDescription), v...))
	})
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Repository {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Repository(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDescription), v...))
	})
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescription), v))
	})
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescription), v))
	})
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescription), v))
	})
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescription), v))
	})
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescription), v))
	})
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescription), v))
	})
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescription), v))
	})
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDescription)))
	})
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDescription)))
	})
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescription), v))
	})
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescription), v))
	})
}

// TotalPrEQ applies the EQ predicate on the "total_pr" field.
func TotalPrEQ(v int64) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTotalPr), v))
	})
}

// TotalPrNEQ applies the NEQ predicate on the "total_pr" field.
func TotalPrNEQ(v int64) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTotalPr), v))
	})
}

// TotalPrIn applies the In predicate on the "total_pr" field.
func TotalPrIn(vs ...int64) predicate.Repository {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Repository(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTotalPr), v...))
	})
}

// TotalPrNotIn applies the NotIn predicate on the "total_pr" field.
func TotalPrNotIn(vs ...int64) predicate.Repository {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Repository(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTotalPr), v...))
	})
}

// TotalPrGT applies the GT predicate on the "total_pr" field.
func TotalPrGT(v int64) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTotalPr), v))
	})
}

// TotalPrGTE applies the GTE predicate on the "total_pr" field.
func TotalPrGTE(v int64) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTotalPr), v))
	})
}

// TotalPrLT applies the LT predicate on the "total_pr" field.
func TotalPrLT(v int64) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTotalPr), v))
	})
}

// TotalPrLTE applies the LTE predicate on the "total_pr" field.
func TotalPrLTE(v int64) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTotalPr), v))
	})
}

// TotalIssueEQ applies the EQ predicate on the "total_issue" field.
func TotalIssueEQ(v int64) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTotalIssue), v))
	})
}

// TotalIssueNEQ applies the NEQ predicate on the "total_issue" field.
func TotalIssueNEQ(v int64) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTotalIssue), v))
	})
}

// TotalIssueIn applies the In predicate on the "total_issue" field.
func TotalIssueIn(vs ...int64) predicate.Repository {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Repository(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTotalIssue), v...))
	})
}

// TotalIssueNotIn applies the NotIn predicate on the "total_issue" field.
func TotalIssueNotIn(vs ...int64) predicate.Repository {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Repository(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTotalIssue), v...))
	})
}

// TotalIssueGT applies the GT predicate on the "total_issue" field.
func TotalIssueGT(v int64) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTotalIssue), v))
	})
}

// TotalIssueGTE applies the GTE predicate on the "total_issue" field.
func TotalIssueGTE(v int64) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTotalIssue), v))
	})
}

// TotalIssueLT applies the LT predicate on the "total_issue" field.
func TotalIssueLT(v int64) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTotalIssue), v))
	})
}

// TotalIssueLTE applies the LTE predicate on the "total_issue" field.
func TotalIssueLTE(v int64) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTotalIssue), v))
	})
}

// GetPullRequestEQ applies the EQ predicate on the "get_pull_request" field.
func GetPullRequestEQ(v bool) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGetPullRequest), v))
	})
}

// GetPullRequestNEQ applies the NEQ predicate on the "get_pull_request" field.
func GetPullRequestNEQ(v bool) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGetPullRequest), v))
	})
}

// GetIssueEQ applies the EQ predicate on the "get_issue" field.
func GetIssueEQ(v bool) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGetIssue), v))
	})
}

// GetIssueNEQ applies the NEQ predicate on the "get_issue" field.
func GetIssueNEQ(v bool) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGetIssue), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Repository {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Repository(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Repository {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Repository(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Repository {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Repository(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Repository {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Repository(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// PushedAtEQ applies the EQ predicate on the "pushed_at" field.
func PushedAtEQ(v time.Time) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPushedAt), v))
	})
}

// PushedAtNEQ applies the NEQ predicate on the "pushed_at" field.
func PushedAtNEQ(v time.Time) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPushedAt), v))
	})
}

// PushedAtIn applies the In predicate on the "pushed_at" field.
func PushedAtIn(vs ...time.Time) predicate.Repository {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Repository(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPushedAt), v...))
	})
}

// PushedAtNotIn applies the NotIn predicate on the "pushed_at" field.
func PushedAtNotIn(vs ...time.Time) predicate.Repository {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Repository(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPushedAt), v...))
	})
}

// PushedAtGT applies the GT predicate on the "pushed_at" field.
func PushedAtGT(v time.Time) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPushedAt), v))
	})
}

// PushedAtGTE applies the GTE predicate on the "pushed_at" field.
func PushedAtGTE(v time.Time) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPushedAt), v))
	})
}

// PushedAtLT applies the LT predicate on the "pushed_at" field.
func PushedAtLT(v time.Time) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPushedAt), v))
	})
}

// PushedAtLTE applies the LTE predicate on the "pushed_at" field.
func PushedAtLTE(v time.Time) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPushedAt), v))
	})
}

// HasPullRequests applies the HasEdge predicate on the "pull_requests" edge.
func HasPullRequests() predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PullRequestsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PullRequestsTable, PullRequestsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPullRequestsWith applies the HasEdge predicate on the "pull_requests" edge with a given conditions (other predicates).
func HasPullRequestsWith(preds ...predicate.PullRequest) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PullRequestsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PullRequestsTable, PullRequestsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasIssues applies the HasEdge predicate on the "issues" edge.
func HasIssues() predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(IssuesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, IssuesTable, IssuesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasIssuesWith applies the HasEdge predicate on the "issues" edge with a given conditions (other predicates).
func HasIssuesWith(preds ...predicate.Issue) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(IssuesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, IssuesTable, IssuesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Repository) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Repository) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
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
func Not(p predicate.Repository) predicate.Repository {
	return predicate.Repository(func(s *sql.Selector) {
		p(s.Not())
	})
}
