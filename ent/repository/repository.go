// Code generated by entc, DO NOT EDIT.

package repository

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the repository type in the database.
	Label = "repository"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldGithubID holds the string denoting the github_id field in the database.
	FieldGithubID = "github_id"
	// FieldOwner holds the string denoting the owner field in the database.
	FieldOwner = "owner"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldTotalPr holds the string denoting the total_pr field in the database.
	FieldTotalPr = "total_pr"
	// FieldTotalIssue holds the string denoting the total_issue field in the database.
	FieldTotalIssue = "total_issue"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldPushedAt holds the string denoting the pushed_at field in the database.
	FieldPushedAt = "pushed_at"
	// EdgePullRequests holds the string denoting the pull_requests edge name in mutations.
	EdgePullRequests = "pull_requests"
	// EdgeIssues holds the string denoting the issues edge name in mutations.
	EdgeIssues = "issues"
	// Table holds the table name of the repository in the database.
	Table = "repositories"
	// PullRequestsTable is the table that holds the pull_requests relation/edge.
	PullRequestsTable = "pull_requests"
	// PullRequestsInverseTable is the table name for the PullRequest entity.
	// It exists in this package in order to avoid circular dependency with the "pullrequest" package.
	PullRequestsInverseTable = "pull_requests"
	// PullRequestsColumn is the table column denoting the pull_requests relation/edge.
	PullRequestsColumn = "repository_id"
	// IssuesTable is the table that holds the issues relation/edge.
	IssuesTable = "issues"
	// IssuesInverseTable is the table name for the Issue entity.
	// It exists in this package in order to avoid circular dependency with the "issue" package.
	IssuesInverseTable = "issues"
	// IssuesColumn is the table column denoting the issues relation/edge.
	IssuesColumn = "repository_id"
)

// Columns holds all SQL columns for repository fields.
var Columns = []string{
	FieldID,
	FieldGithubID,
	FieldOwner,
	FieldName,
	FieldDescription,
	FieldTotalPr,
	FieldTotalIssue,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldPushedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// GithubIDValidator is a validator for the "github_id" field. It is called by the builders before save.
	GithubIDValidator func(string) error
	// OwnerValidator is a validator for the "owner" field. It is called by the builders before save.
	OwnerValidator func(string) error
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultTotalPr holds the default value on creation for the "total_pr" field.
	DefaultTotalPr int64
	// TotalPrValidator is a validator for the "total_pr" field. It is called by the builders before save.
	TotalPrValidator func(int64) error
	// DefaultTotalIssue holds the default value on creation for the "total_issue" field.
	DefaultTotalIssue int64
	// TotalIssueValidator is a validator for the "total_issue" field. It is called by the builders before save.
	TotalIssueValidator func(int64) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// DefaultPushedAt holds the default value on creation for the "pushed_at" field.
	DefaultPushedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
