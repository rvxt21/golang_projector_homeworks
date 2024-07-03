package main

import (
	"main/process"
	"regexp"
)

func main() {
	pattern := regexp.MustCompile(`\(?\d{3}\)?[-.\s]?\d{3}[-.\s]?\d{4}`)

	filePath := "files/numbers.txt"

	fileContent := process.ReadFile(filePath)

	process.FindMatches(pattern, fileContent)
}
