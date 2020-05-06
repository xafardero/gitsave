package github

import (
	"context"

	"github.com/google/go-github/v28/github"
	"golang.org/x/oauth2"
)

func getAllGithubRepositories(accessToken, organization string) []*github.Repository {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	opt := &github.RepositoryListByOrgOptions{
		Type:        "all",
		ListOptions: github.ListOptions{PerPage: 150},
	}

	// list all repositories for the authenticated user
	repos, _, _ := client.Repositories.ListByOrg(ctx, organization, opt)

	return repos
}

func transformGhReposToRepos(ghRepos []*github.Repository) []Repository {
	repositories := []Repository{}

	for _, repo := range ghRepos {
		r := Repository{
			ID:          repo.GetID(),
			Name:        repo.GetName(),
			Language:    repo.GetLanguage(),
			Description: repo.GetDescription(),
		}
		repositories = append(repositories, r)
	}

	return repositories
}
