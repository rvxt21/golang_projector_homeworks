package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/rs/zerolog/log"
)

func main() {
	// `[А-ЯҐЄІЇ]+[А-ЩЬЮЯҐЄІЇа-щьюяґєії]+` "Така", "Київ", "Отож"
	//`[А-ЩЬЮЯҐЄІЇа-щьюяґєії]+,` "душу,""ні"
	//^[А-ЯҐЄІЇ][а-щьюяґєії]+

	pattern, err := regexp.Compile(`[А-ЩЬЮЯҐЄІЇа-щьюяґєії]+,`)
	if err != nil {
		panic(err.Error())
	}
	filePath := "files/text.txt"

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to open file")
	}
	matches := pattern.FindAllString(string(fileContent), -1)

	for i, v := range matches {
		fmt.Println(i, v)
	}
}
