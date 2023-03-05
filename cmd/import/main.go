package main

import (
	"fmt"
	"os"

	"github.com/shinnosuke-K/github-dev-insight/pkg/adapter"
	"github.com/shinnosuke-K/github-dev-insight/pkg/app/service"
	"github.com/shinnosuke-K/github-dev-insight/pkg/env"

	"github.com/spf13/cobra"
)

func init() {
	if err := env.Load(); err != nil {
		panic(fmt.Sprintf("failed to load env config. %s", err))
	}
}

func main() {
	if err := _main(); err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func _main() error {
	ad, err := adapter.NewAdapter()
	if err != nil {
		return fmt.Errorf("faild to create adapter. %w", err)
	}

	var (
		commitSvc   = service.NewCommitService(ad)
		issueSvc    = service.NewIssueService(ad)
		prSvc       = service.NewPullRequestService(ad)
		repoSvc     = service.NewRepositoryService(ad)
		rootCmd     = &cobra.Command{}
		subCommands = []*cobra.Command{
			{
				Use:  "commit",
				RunE: func(cmd *cobra.Command, args []string) error { return commitSvc.ImportCommit(cmd.Context()) },
			},
			{
				Use:  "issue",
				RunE: func(cmd *cobra.Command, args []string) error { return issueSvc.ImportIssue(cmd.Context()) },
			},
			{
				Use:  "pr",
				RunE: func(cmd *cobra.Command, args []string) error { return prSvc.ImportPullRequest(cmd.Context()) },
			},
			{
				Use:  "repo",
				RunE: func(cmd *cobra.Command, args []string) error { return repoSvc.ImportRepositories(cmd.Context()) },
			},
		}
	)
	rootCmd.AddCommand(subCommands...)
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
