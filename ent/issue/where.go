// Code generated by entc, DO NOT EDIT.

package issue

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/shinnosuke-K/github-dev-insight/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
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
func IDNotIn(ids ...uuid.UUID) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
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
func IDGT(id uuid.UUID) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// GithubID applies equality check predicate on the "github_id" field. It's identical to GithubIDEQ.
func GithubID(v string) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGithubID), v))
	})
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// LastEditedAt applies equality check predicate on the "last_edited_at" field. It's identical to LastEditedAtEQ.
func LastEditedAt(v time.Time) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastEditedAt), v))
	})
}

// ClosedAt applies equality check predicate on the "closed_at" field. It's identical to ClosedAtEQ.
func ClosedAt(v time.Time) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldClosedAt), v))
	})
}

// GithubIDEQ applies the EQ predicate on the "github_id" field.
func GithubIDEQ(v string) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGithubID), v))
	})
}

// GithubIDNEQ applies the NEQ predicate on the "github_id" field.
func GithubIDNEQ(v string) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGithubID), v))
	})
}

// GithubIDIn applies the In predicate on the "github_id" field.
func GithubIDIn(vs ...string) predicate.Issue {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Issue(func(s *sql.Selector) {
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
func GithubIDNotIn(vs ...string) predicate.Issue {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Issue(func(s *sql.Selector) {
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
func GithubIDGT(v string) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGithubID), v))
	})
}

// GithubIDGTE applies the GTE predicate on the "github_id" field.
func GithubIDGTE(v string) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGithubID), v))
	})
}

// GithubIDLT applies the LT predicate on the "github_id" field.
func GithubIDLT(v string) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGithubID), v))
	})
}

// GithubIDLTE applies the LTE predicate on the "github_id" field.
func GithubIDLTE(v string) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGithubID), v))
	})
}

// GithubIDContains applies the Contains predicate on the "github_id" field.
func GithubIDContains(v string) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldGithubID), v))
	})
}

// GithubIDHasPrefix applies the HasPrefix predicate on the "github_id" field.
func GithubIDHasPrefix(v string) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldGithubID), v))
	})
}

// GithubIDHasSuffix applies the HasSuffix predicate on the "github_id" field.
func GithubIDHasSuffix(v string) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldGithubID), v))
	})
}

// GithubIDEqualFold applies the EqualFold predicate on the "github_id" field.
func GithubIDEqualFold(v string) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldGithubID), v))
	})
}

// GithubIDContainsFold applies the ContainsFold predicate on the "github_id" field.
func GithubIDContainsFold(v string) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldGithubID), v))
	})
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitle), v))
	})
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Issue {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Issue(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTitle), v...))
	})
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Issue {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Issue(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTitle), v...))
	})
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitle), v))
	})
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitle), v))
	})
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitle), v))
	})
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitle), v))
	})
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitle), v))
	})
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitle), v))
	})
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitle), v))
	})
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitle), v))
	})
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitle), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Issue {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Issue(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.Issue {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Issue(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Issue {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Issue(func(s *sql.Selector) {
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
func UpdatedAtNotIn(vs ...time.Time) predicate.Issue {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Issue(func(s *sql.Selector) {
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
func UpdatedAtGT(v time.Time) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// LastEditedAtEQ applies the EQ predicate on the "last_edited_at" field.
func LastEditedAtEQ(v time.Time) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastEditedAt), v))
	})
}

// LastEditedAtNEQ applies the NEQ predicate on the "last_edited_at" field.
func LastEditedAtNEQ(v time.Time) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLastEditedAt), v))
	})
}

// LastEditedAtIn applies the In predicate on the "last_edited_at" field.
func LastEditedAtIn(vs ...time.Time) predicate.Issue {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Issue(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLastEditedAt), v...))
	})
}

// LastEditedAtNotIn applies the NotIn predicate on the "last_edited_at" field.
func LastEditedAtNotIn(vs ...time.Time) predicate.Issue {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Issue(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLastEditedAt), v...))
	})
}

// LastEditedAtGT applies the GT predicate on the "last_edited_at" field.
func LastEditedAtGT(v time.Time) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLastEditedAt), v))
	})
}

// LastEditedAtGTE applies the GTE predicate on the "last_edited_at" field.
func LastEditedAtGTE(v time.Time) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLastEditedAt), v))
	})
}

// LastEditedAtLT applies the LT predicate on the "last_edited_at" field.
func LastEditedAtLT(v time.Time) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLastEditedAt), v))
	})
}

// LastEditedAtLTE applies the LTE predicate on the "last_edited_at" field.
func LastEditedAtLTE(v time.Time) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLastEditedAt), v))
	})
}

// ClosedAtEQ applies the EQ predicate on the "closed_at" field.
func ClosedAtEQ(v time.Time) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldClosedAt), v))
	})
}

// ClosedAtNEQ applies the NEQ predicate on the "closed_at" field.
func ClosedAtNEQ(v time.Time) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldClosedAt), v))
	})
}

// ClosedAtIn applies the In predicate on the "closed_at" field.
func ClosedAtIn(vs ...time.Time) predicate.Issue {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Issue(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldClosedAt), v...))
	})
}

// ClosedAtNotIn applies the NotIn predicate on the "closed_at" field.
func ClosedAtNotIn(vs ...time.Time) predicate.Issue {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Issue(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldClosedAt), v...))
	})
}

// ClosedAtGT applies the GT predicate on the "closed_at" field.
func ClosedAtGT(v time.Time) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldClosedAt), v))
	})
}

// ClosedAtGTE applies the GTE predicate on the "closed_at" field.
func ClosedAtGTE(v time.Time) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldClosedAt), v))
	})
}

// ClosedAtLT applies the LT predicate on the "closed_at" field.
func ClosedAtLT(v time.Time) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldClosedAt), v))
	})
}

// ClosedAtLTE applies the LTE predicate on the "closed_at" field.
func ClosedAtLTE(v time.Time) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldClosedAt), v))
	})
}

// HasRepository applies the HasEdge predicate on the "repository" edge.
func HasRepository() predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RepositoryTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RepositoryTable, RepositoryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRepositoryWith applies the HasEdge predicate on the "repository" edge with a given conditions (other predicates).
func HasRepositoryWith(preds ...predicate.Repository) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
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
func And(predicates ...predicate.Issue) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Issue) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
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
func Not(p predicate.Issue) predicate.Issue {
	return predicate.Issue(func(s *sql.Selector) {
		p(s.Not())
	})
}
