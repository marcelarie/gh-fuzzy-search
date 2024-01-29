package gh

import (
	"context"
	"fmt"
	"github.com/google/go-github/v58/github"
	"os"
)

const (
	GithubTokenEnvVar  = "GITHUB_TOKEN"
	UserNotFoundError  = "Error: User %s not found"
	TokenNotFoundError = "Error: You must set GITHUB_TOKEN environment variable"
)

type ErrUserNotFound struct {
	Username string
}

func (e *ErrUserNotFound) Error() string {
	return fmt.Sprintf(UserNotFoundError, e.Username)
}

func NewErrUserNotFound(username string) error {
	return &ErrUserNotFound{Username: username}
}

func GetGithubToken() (string, error) {
	ghToken := os.Getenv(GithubTokenEnvVar)

	if ghToken == "" {
		return "", fmt.Errorf(TokenNotFoundError)
	}

	return ghToken, nil
}

func GetRepos(username string) ([]string, error) {
	ghToken, err := GetGithubToken()

	if err != nil {
		return []string{}, err
	}

	ctx := context.Background()
	client := github.NewClient(nil).WithAuthToken(ghToken)

	opts := &github.RepositoryListByUserOptions{
		Type:        "owner",
		Sort:        "updated",
		ListOptions: github.ListOptions{PerPage: 100},
	}

	var repoNames []string
	for {
		repos, resp, err := client.Repositories.ListByUser(ctx, username, opts)

		if err != nil {
			if resp != nil && resp.StatusCode == 404 {
				return []string{}, NewErrUserNotFound(username)
			}
			return []string{}, err
		}

		for _, repo := range repos {
			repoNames = append(repoNames, *repo.Name)
		}
		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	return repoNames, nil
}

func GetUsers(search string) ([]string, error) {
	ghToken, err := GetGithubToken()
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	client := github.NewClient(nil).WithAuthToken(ghToken)
	opts := &github.SearchOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}
	var usernames []string
	for {
		users, resp, err := client.Search.Users(ctx, search, opts)
		if err != nil {
			return nil, err
		}
		for _, user := range users.Users {
			usernames = append(usernames, *user.Login)
		}
		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}
	return usernames, nil
}
