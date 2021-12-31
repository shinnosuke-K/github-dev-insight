// Code generated by entc, DO NOT EDIT.

package commits

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/shinnosuke-K/github-dev-insight/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
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
func IDNotIn(ids ...uuid.UUID) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
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
func IDGT(id uuid.UUID) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// GithubID applies equality check predicate on the "github_id" field. It's identical to GithubIDEQ.
func GithubID(v string) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGithubID), v))
	})
}

// Message applies equality check predicate on the "message" field. It's identical to MessageEQ.
func Message(v string) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMessage), v))
	})
}

// CommittedAt applies equality check predicate on the "committed_at" field. It's identical to CommittedAtEQ.
func CommittedAt(v time.Time) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCommittedAt), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// GithubIDEQ applies the EQ predicate on the "github_id" field.
func GithubIDEQ(v string) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGithubID), v))
	})
}

// GithubIDNEQ applies the NEQ predicate on the "github_id" field.
func GithubIDNEQ(v string) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGithubID), v))
	})
}

// GithubIDIn applies the In predicate on the "github_id" field.
func GithubIDIn(vs ...string) predicate.Commits {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Commits(func(s *sql.Selector) {
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
func GithubIDNotIn(vs ...string) predicate.Commits {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Commits(func(s *sql.Selector) {
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
func GithubIDGT(v string) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGithubID), v))
	})
}

// GithubIDGTE applies the GTE predicate on the "github_id" field.
func GithubIDGTE(v string) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGithubID), v))
	})
}

// GithubIDLT applies the LT predicate on the "github_id" field.
func GithubIDLT(v string) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGithubID), v))
	})
}

// GithubIDLTE applies the LTE predicate on the "github_id" field.
func GithubIDLTE(v string) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGithubID), v))
	})
}

// GithubIDContains applies the Contains predicate on the "github_id" field.
func GithubIDContains(v string) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldGithubID), v))
	})
}

// GithubIDHasPrefix applies the HasPrefix predicate on the "github_id" field.
func GithubIDHasPrefix(v string) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldGithubID), v))
	})
}

// GithubIDHasSuffix applies the HasSuffix predicate on the "github_id" field.
func GithubIDHasSuffix(v string) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldGithubID), v))
	})
}

// GithubIDEqualFold applies the EqualFold predicate on the "github_id" field.
func GithubIDEqualFold(v string) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldGithubID), v))
	})
}

// GithubIDContainsFold applies the ContainsFold predicate on the "github_id" field.
func GithubIDContainsFold(v string) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldGithubID), v))
	})
}

// MessageEQ applies the EQ predicate on the "message" field.
func MessageEQ(v string) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMessage), v))
	})
}

// MessageNEQ applies the NEQ predicate on the "message" field.
func MessageNEQ(v string) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMessage), v))
	})
}

// MessageIn applies the In predicate on the "message" field.
func MessageIn(vs ...string) predicate.Commits {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Commits(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMessage), v...))
	})
}

// MessageNotIn applies the NotIn predicate on the "message" field.
func MessageNotIn(vs ...string) predicate.Commits {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Commits(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMessage), v...))
	})
}

// MessageGT applies the GT predicate on the "message" field.
func MessageGT(v string) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMessage), v))
	})
}

// MessageGTE applies the GTE predicate on the "message" field.
func MessageGTE(v string) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMessage), v))
	})
}

// MessageLT applies the LT predicate on the "message" field.
func MessageLT(v string) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMessage), v))
	})
}

// MessageLTE applies the LTE predicate on the "message" field.
func MessageLTE(v string) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMessage), v))
	})
}

// MessageContains applies the Contains predicate on the "message" field.
func MessageContains(v string) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMessage), v))
	})
}

// MessageHasPrefix applies the HasPrefix predicate on the "message" field.
func MessageHasPrefix(v string) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMessage), v))
	})
}

// MessageHasSuffix applies the HasSuffix predicate on the "message" field.
func MessageHasSuffix(v string) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMessage), v))
	})
}

// MessageEqualFold applies the EqualFold predicate on the "message" field.
func MessageEqualFold(v string) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMessage), v))
	})
}

// MessageContainsFold applies the ContainsFold predicate on the "message" field.
func MessageContainsFold(v string) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMessage), v))
	})
}

// CommittedAtEQ applies the EQ predicate on the "committed_at" field.
func CommittedAtEQ(v time.Time) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCommittedAt), v))
	})
}

// CommittedAtNEQ applies the NEQ predicate on the "committed_at" field.
func CommittedAtNEQ(v time.Time) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCommittedAt), v))
	})
}

// CommittedAtIn applies the In predicate on the "committed_at" field.
func CommittedAtIn(vs ...time.Time) predicate.Commits {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Commits(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCommittedAt), v...))
	})
}

// CommittedAtNotIn applies the NotIn predicate on the "committed_at" field.
func CommittedAtNotIn(vs ...time.Time) predicate.Commits {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Commits(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCommittedAt), v...))
	})
}

// CommittedAtGT applies the GT predicate on the "committed_at" field.
func CommittedAtGT(v time.Time) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCommittedAt), v))
	})
}

// CommittedAtGTE applies the GTE predicate on the "committed_at" field.
func CommittedAtGTE(v time.Time) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCommittedAt), v))
	})
}

// CommittedAtLT applies the LT predicate on the "committed_at" field.
func CommittedAtLT(v time.Time) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCommittedAt), v))
	})
}

// CommittedAtLTE applies the LTE predicate on the "committed_at" field.
func CommittedAtLTE(v time.Time) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCommittedAt), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Commits {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Commits(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.Commits {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Commits(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// HasPullRequest applies the HasEdge predicate on the "pull_request" edge.
func HasPullRequest() predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PullRequestTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PullRequestTable, PullRequestColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPullRequestWith applies the HasEdge predicate on the "pull_request" edge with a given conditions (other predicates).
func HasPullRequestWith(preds ...predicate.PullRequest) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PullRequestInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PullRequestTable, PullRequestColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Commits) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Commits) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
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
func Not(p predicate.Commits) predicate.Commits {
	return predicate.Commits(func(s *sql.Selector) {
		p(s.Not())
	})
}