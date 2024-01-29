package gh

import (
	"fmt"
	"os"
)

func CheckGithubTokenFromEnv() {
	ghToken := os.Getenv("GITHUB_TOKEN")

	if ghToken == "" {
		fmt.Println("Error: You must set GITHUB_TOKEN environment variable")
		fmt.Println("See https://help.github.com/en/github/authenticating-to-github/creating-a-personal-access-token-for-the-command-line")
		os.Exit(0)
	}
}
