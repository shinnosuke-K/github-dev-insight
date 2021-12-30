// Code generated by entc, DO NOT EDIT.

package commit

import (
	"time"
)

const (
	// Label holds the string label denoting the commit type in the database.
	Label = "commit"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPullrequestID holds the string denoting the pullrequest_id field in the database.
	FieldPullrequestID = "pullrequest_id"
	// FieldGithubID holds the string denoting the github_id field in the database.
	FieldGithubID = "github_id"
	// FieldMessage holds the string denoting the message field in the database.
	FieldMessage = "message"
	// FieldCommittedAt holds the string denoting the committed_at field in the database.
	FieldCommittedAt = "committed_at"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgePullRequests holds the string denoting the pull_requests edge name in mutations.
	EdgePullRequests = "pull_requests"
	// Table holds the table name of the commit in the database.
	Table = "commits"
	// PullRequestsTable is the table that holds the pull_requests relation/edge.
	PullRequestsTable = "commits"
	// PullRequestsInverseTable is the table name for the PullRequest entity.
	// It exists in this package in order to avoid circular dependency with the "pullrequest" package.
	PullRequestsInverseTable = "pull_requests"
	// PullRequestsColumn is the table column denoting the pull_requests relation/edge.
	PullRequestsColumn = "pull_request_commits"
)

// Columns holds all SQL columns for commit fields.
var Columns = []string{
	FieldID,
	FieldPullrequestID,
	FieldGithubID,
	FieldMessage,
	FieldCommittedAt,
	FieldCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "commits"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"pull_request_commits",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// PullrequestIDValidator is a validator for the "pullrequest_id" field. It is called by the builders before save.
	PullrequestIDValidator func(string) error
	// GithubIDValidator is a validator for the "github_id" field. It is called by the builders before save.
	GithubIDValidator func(string) error
	// DefaultCommittedAt holds the default value on creation for the "committed_at" field.
	DefaultCommittedAt func() time.Time
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)
