package entity

type TargetType string

const (
	TargetTypePullRequest TargetType = "pull_request"
	TargetTypeIssue       TargetType = "issue"
	TargetTypeCommit      TargetType = "commit"
)
