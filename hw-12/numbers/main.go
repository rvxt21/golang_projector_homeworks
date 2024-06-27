package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/rs/zerolog/log"
)

func main() {
	pattern, err := regexp.Compile(`\(?\d{3}\)?[-.\s]?\d{3}[-.\s]?\d{4}`)
	if err != nil {
		panic(err.Error())
	}
	filePath := "files/numbers.txt"

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to open file")
	}
	matches := pattern.FindAllString(string(fileContent), -1)

	for i, v := range matches {
		fmt.Println(i, v)
	}
}
