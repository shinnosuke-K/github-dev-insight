package entity

type TargetType string

const (
	PR     TargetType = "pull_request"
	Issue  TargetType = "issue"
	Commit TargetType = "commit"
)
