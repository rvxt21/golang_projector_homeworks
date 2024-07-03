package process

import (
	"fmt"
	"os"
	"regexp"

	"github.com/rs/zerolog/log"
)

func CompileRegex(pattern string) *regexp.Regexp {
	return regexp.MustCompile(pattern)
}

func ReadFile(filePath string) string {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to open file")
	}
	return string(fileContent)
}

func FindMatches(pattern *regexp.Regexp, content string) {
	matches := pattern.FindAllString(content, -1)

	for i, v := range matches {
		fmt.Println(i, v)
	}
}
