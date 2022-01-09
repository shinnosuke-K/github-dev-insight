// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CommitsColumns holds the columns for the "commits" table.
	CommitsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "github_id", Type: field.TypeString, Size: 255},
		{Name: "message", Type: field.TypeString, Size: 2147483647},
		{Name: "committed_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "pull_request_id", Type: field.TypeUUID, Nullable: true},
	}
	// CommitsTable holds the schema information for the "commits" table.
	CommitsTable = &schema.Table{
		Name:       "commits",
		Columns:    CommitsColumns,
		PrimaryKey: []*schema.Column{CommitsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "commits_pull_requests_commits",
				Columns:    []*schema.Column{CommitsColumns[5]},
				RefColumns: []*schema.Column{PullRequestsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// IssuesColumns holds the columns for the "issues" table.
	IssuesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "github_id", Type: field.TypeString, Size: 255},
		{Name: "title", Type: field.TypeString, Size: 255},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "last_edited_at", Type: field.TypeTime, Nullable: true},
		{Name: "closed_at", Type: field.TypeTime, Nullable: true},
		{Name: "repository_id", Type: field.TypeUUID, Nullable: true},
	}
	// IssuesTable holds the schema information for the "issues" table.
	IssuesTable = &schema.Table{
		Name:       "issues",
		Columns:    IssuesColumns,
		PrimaryKey: []*schema.Column{IssuesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "issues_repositories_issues",
				Columns:    []*schema.Column{IssuesColumns[7]},
				RefColumns: []*schema.Column{RepositoriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PullRequestsColumns holds the columns for the "pull_requests" table.
	PullRequestsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "github_id", Type: field.TypeString, Size: 255},
		{Name: "title", Type: field.TypeString, Size: 255},
		{Name: "total_commits", Type: field.TypeInt64, Default: 0},
		{Name: "get_commit", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "closed_at", Type: field.TypeTime, Nullable: true},
		{Name: "merged_at", Type: field.TypeTime, Nullable: true},
		{Name: "repository_id", Type: field.TypeUUID, Nullable: true},
	}
	// PullRequestsTable holds the schema information for the "pull_requests" table.
	PullRequestsTable = &schema.Table{
		Name:       "pull_requests",
		Columns:    PullRequestsColumns,
		PrimaryKey: []*schema.Column{PullRequestsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "pull_requests_repositories_pull_requests",
				Columns:    []*schema.Column{PullRequestsColumns[9]},
				RefColumns: []*schema.Column{RepositoriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// RepositoriesColumns holds the columns for the "repositories" table.
	RepositoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "github_id", Type: field.TypeString, Size: 255},
		{Name: "owner", Type: field.TypeString, Size: 255},
		{Name: "name", Type: field.TypeString, Size: 255},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "total_pr", Type: field.TypeInt64, Default: 0},
		{Name: "total_issue", Type: field.TypeInt64, Default: 0},
		{Name: "get_pull_request", Type: field.TypeBool, Default: false},
		{Name: "get_issue", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "pushed_at", Type: field.TypeTime},
	}
	// RepositoriesTable holds the schema information for the "repositories" table.
	RepositoriesTable = &schema.Table{
		Name:       "repositories",
		Columns:    RepositoriesColumns,
		PrimaryKey: []*schema.Column{RepositoriesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CommitsTable,
		IssuesTable,
		PullRequestsTable,
		RepositoriesTable,
	}
)

func init() {
	CommitsTable.ForeignKeys[0].RefTable = PullRequestsTable
	IssuesTable.ForeignKeys[0].RefTable = RepositoriesTable
	PullRequestsTable.ForeignKeys[0].RefTable = RepositoriesTable
}
