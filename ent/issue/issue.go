// Code generated by entc, DO NOT EDIT.

package issue

import (
	"time"
)

const (
	// Label holds the string label denoting the issue type in the database.
	Label = "issue"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRepositoryID holds the string denoting the repository_id field in the database.
	FieldRepositoryID = "repository_id"
	// FieldGithubID holds the string denoting the github_id field in the database.
	FieldGithubID = "github_id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldLastEditedAt holds the string denoting the last_edited_at field in the database.
	FieldLastEditedAt = "last_edited_at"
	// FieldClosedAt holds the string denoting the closed_at field in the database.
	FieldClosedAt = "closed_at"
	// EdgeRepository holds the string denoting the repository edge name in mutations.
	EdgeRepository = "repository"
	// Table holds the table name of the issue in the database.
	Table = "issues"
	// RepositoryTable is the table that holds the repository relation/edge.
	RepositoryTable = "issues"
	// RepositoryInverseTable is the table name for the Repository entity.
	// It exists in this package in order to avoid circular dependency with the "repository" package.
	RepositoryInverseTable = "repositories"
	// RepositoryColumn is the table column denoting the repository relation/edge.
	RepositoryColumn = "repository_issues"
)

// Columns holds all SQL columns for issue fields.
var Columns = []string{
	FieldID,
	FieldRepositoryID,
	FieldGithubID,
	FieldTitle,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldLastEditedAt,
	FieldClosedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "issues"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"repository_issues",
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
	// RepositoryIDValidator is a validator for the "repository_id" field. It is called by the builders before save.
	RepositoryIDValidator func(string) error
	// GithubIDValidator is a validator for the "github_id" field. It is called by the builders before save.
	GithubIDValidator func(string) error
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// DefaultLastEditedAt holds the default value on creation for the "last_edited_at" field.
	DefaultLastEditedAt func() time.Time
	// DefaultClosedAt holds the default value on creation for the "closed_at" field.
	DefaultClosedAt func() time.Time
)
