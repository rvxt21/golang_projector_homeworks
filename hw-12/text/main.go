package main

import (
	"main/process"
)

func main() {
	// `[А-ЯҐЄІЇ]+[А-ЩЬЮЯҐЄІЇа-щьюяґєії]+` "Така", "Київ", "Отож"
	//`[А-ЩЬЮЯҐЄІЇа-щьюяґєії]+,` "душу,""ні"
	//^[А-ЯҐЄІЇ][а-щьюяґєії]+

	pattern := process.CompileRegex(`[А-ЩЬЮЯҐЄІЇа-щьюяґєії]+,`)
	filePath := "files/text.txt"

	fileContent := process.ReadFile(filePath)

	process.FindMatches(pattern, fileContent)
}
