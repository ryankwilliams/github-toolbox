package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {
	var githubAccessToken, githubOrganization, githubRepository string

	if githubAccessToken = os.Getenv("GITHUB_API_TOKEN"); githubAccessToken == "" {
		log.Fatal("GITHUB_API_TOKEN undefined.")
	}

	if githubOrganization = os.Getenv("GITHUB_ORGANIZATION"); githubOrganization == "" {
		log.Fatal("GITHUB_ORGANIZATION undefined.")
	}

	if githubRepository = os.Getenv("GITHUB_REPOSITORY"); githubRepository == "" {
		log.Fatal("GITHUB_REPOSITORY undefined.")
	}

	githubUsername := "ryankwilliams"
	prState := "closed"

	ctx := context.Background()
	token := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: githubAccessToken,
		},
	)

	client := github.NewClient(oauth2.NewClient(ctx, token))

	prs, _, err := client.PullRequests.List(
		ctx,
		githubOrganization,
		githubRepository,
		&github.PullRequestListOptions{
			State: prState,
			ListOptions: github.ListOptions{
				PerPage: 50,
			},
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s/%s", githubOrganization, githubRepository)
	for _, pr := range prs {
		if *pr.User.Login == githubUsername {
			fmt.Printf(`  #%d
    Title   : %s
    URL     : %s
    Created : %s
    Closed  : %s`,
				*pr.Number,
				*pr.Title,
				*pr.HTMLURL,
				pr.GetCreatedAt(),
				pr.GetClosedAt())
			fmt.Println()
		}
	}
}
