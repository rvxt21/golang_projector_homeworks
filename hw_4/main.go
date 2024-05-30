package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Please input string to find:")
	var UserInputToFindSlice string
	fmt.Scan(&UserInputToFindSlice)
	fmt.Printf("Your word to find: %s.\n", UserInputToFindSlice)
	CreatedTextByUser := CreateText()
	FindSliceInText(CreatedTextByUser, UserInputToFindSlice)
}

func FindSliceInText(Slice []string, WordToFind string) {
	fmt.Println("******************************")
	fmt.Println("Results:")
	found := false
	WordToFind = strings.ToLower(WordToFind)
	for i, val := range Slice {
		if strings.Contains(strings.ToLower(val), WordToFind) {
			fmt.Printf("Line %d: %s\n", i+1, val)
			found = true
		}
	}
	if !found {
		fmt.Println("No matches found.")
	}
}

func CreateText() []string {
	var UserCreatedText []string
	var UserInput string
	fmt.Println("******************************")
	fmt.Println("Input words, for new word press Enter. To stop press 'q'")
	for true {
		fmt.Print(">")
		fmt.Scan(&UserInput)
		if UserInput == "q" {
			break
		}
		UserCreatedText = append(UserCreatedText, UserInput)
	}
	fmt.Println("******************************")
	fmt.Println(UserCreatedText)
	return UserCreatedText
}
