package main

import (
	"bufio"
	"fmt"
	"gh-fuzzy-search/gh"
	"os"
	"strings"
)

const usageMessage = `Usage gh-fuzzy-search [arguments]
Arguments:
  -h, --help: Show this help message
  -u, --user: Search for a GitHub user
  -v, --version: Show the version of this program`

func printColoredText(text string, color Color) {
	fmt.Println(string(color) + text + "\033[0m")
}

func readUsername() string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter search string (Ctrl+C to exit)")
		fmt.Print(">> ")
		search, _ := reader.ReadString('\n')
		search = strings.TrimSpace(search) // remove newline character

		usernames, err := gh.GetUsers(search)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Println("Usernames:", usernames)
	}
}

func userCommand(args []string) {
	if len(args) == 1 {
		// fmt.Println("Error: You must pass at least a GitHub username")
		// fmt.Println("Example: gh-fuzzy-search -u <username>")
		// os.Exit(0)
		readUsername()
	}

	username := args[1]

	// repos, err := gh.GetRepos(username)
	users, err := gh.GetUsers(username)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, user := range users {
		fmt.Println(user)
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
		userCommand(args)
		os.Exit(0)
	case "-r", "--repo":
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
