package main

import (
	"fmt"
	"os"
)

const usageMessage = `Usage gh-fuzzy-search [arguments]
Arguments:
  -h, --help: Show this help message
  -v, --version: Show the version of this program`

func printColoredText(text string, color Color) {
	fmt.Println(string(color) + text + "\033[0m")
}

func userCommand(args []string) {
	if len(args) == 1 {
		fmt.Println("Error: You must pass at least a GitHub username")
		fmt.Println("Example: gh-fuzzy-search -u <username>")
		os.Exit(0)
	}

	fmt.Println("user: " + args[1])
}

func checkGithubTokenFromEnv() {
	ghToken := os.Getenv("GITHUB_TOKEN")
	if ghToken == "" {
		fmt.Println("Error: You must set GITHUB_TOKEN environment variable")
		fmt.Println("See https://help.github.com/en/github/authenticating-to-github/creating-a-personal-access-token-for-the-command-line")
		os.Exit(0)
	}
}

func checkArgs(args []string) {
	if len(args) == 0 {
		printColoredText("Error: You must pass at least one argument", ColorRed)
		printColoredText(usageMessage, ColorWhite)
		os.Exit(0)
	}

	firstArg := args[0]

	switch firstArg {
	case "-h", "--help":
		printColoredText(usageMessage, ColorWhite)
		os.Exit(0)
	case "-v", "--version":
		const versionMessage = "gh-fuzzy-search version 0.0.1"
		printColoredText(versionMessage, ColorWhite)
		os.Exit(0)
	case "-u", "--user":
		checkGithubTokenFromEnv()
		userCommand(args)
		os.Exit(0)
	case "-r", "--repo":
		checkGithubTokenFromEnv()
		// TODO: Implement this
		printColoredText("repo", ColorWhite)
		os.Exit(0)
	default:
		printColoredText("Error: Invalid argument", ColorRed)
		os.Exit(0)
	}
}

func main() {
	args := os.Args[1:]

	checkArgs(args)
}
