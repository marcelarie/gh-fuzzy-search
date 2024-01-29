package main

import (
	"fmt"
	"os"
)

const usage = `Usage gh-fuzzy-search [arguments]
Arguments:
  -h, --help: Show this help message
  -v, --version: Show the version of this program`
const version = "gh-fuzzy-search version 0.0.1"

type Color string

const (
	ColorRed    Color = "\033[31m"
	ColorGreen        = "\033[32m"
	ColorBlue         = "\033[34m"
	ColorYellow       = "\033[33m"
	ColorPurple       = "\033[35m"
	ColorCyan         = "\033[36m"
	ColorWhite        = "\033[37m"
)

func printColoredText(text string, color Color) {
	fmt.Println(string(color) + text + "\033[0m")
}

func checkArgs(args []string) {
	for _, arg := range args {
		switch arg {
		case "-h", "--help":
			printColoredText(usage, ColorWhite)
			os.Exit(0)
		case "-v", "--version":
			printColoredText(version, ColorWhite)
			os.Exit(0)
		default:
			printColoredText("Error: Invalid argument", ColorRed)
			os.Exit(0)
		}
	}
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		printColoredText("Error: You must pass at least one argument", ColorRed)
		printColoredText(usage, ColorWhite)
		os.Exit(0)
	}

	checkArgs(args)
}
