package gh

import (
	"context"
	"fmt"
	"github.com/google/go-github/v58/github"
	"os"
)

// ErrUserNotFound is an error indicating that the GitHub user was not found.
type ErrUserNotFound struct {
	Username string
}

func (e *ErrUserNotFound) Error() string {
	return fmt.Sprintf("Error: User %s not found", e.Username)
}

// NewErrUserNotFound creates a new instance of ErrUserNotFound.
func NewErrUserNotFound(username string) error {
	return &ErrUserNotFound{Username: username}
}

func GetGithubToken() (string, error) {
	ghToken := os.Getenv("GITHUB_TOKEN")

	if ghToken == "" {
		return "", fmt.Errorf("Error: You must set GITHUB_TOKEN environment variable")
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

	opts := &github.RepositoryListByUserOptions{Type: "owner", Sort: "updated", ListOptions: github.ListOptions{PerPage: 100}}
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
