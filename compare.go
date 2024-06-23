package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v46/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
	"log"
	"os"
)

var compareCmd = &cobra.Command{
	Use:   "compare",
	Short: "Compare branch with master and check coding principles",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			log.Fatal("Usage: compare [repo_owner] [repo_name] [branch]")
		}
		repoOwner, repoName, branch := args[0], args[1], args[2]
		checkChanges(repoOwner, repoName, branch)
	},
}

func init() {
	rootCmd.AddCommand(compareCmd)
}
func checkChanges(owner string, repo string, branch string) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	compare, _, err := client.Repositories.CompareCommits(ctx, owner, repo, "master", branch, nil)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range compare.Files {
		fmt.Printf("Checking file: %s\n", *file.Filename)
		content := *file.Patch
		checkCodingPrinciples(content)
	}
}
