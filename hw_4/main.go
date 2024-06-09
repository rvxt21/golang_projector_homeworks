package main

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"strings"
)

func main() {
	fmt.Println("*******************FIRST****************************")
	fmt.Println("Please input string to find:")
	var UserInputToFindSlice string
	fmt.Scan(&UserInputToFindSlice)
	fmt.Printf("Your word to find: %s.\n", UserInputToFindSlice)
	CreatedTextByUser := CreateText()
	FindSliceInText(CreatedTextByUser, UserInputToFindSlice)

	fmt.Println("*******************SECOND****************************")
	startSlice := CreateSlice()
	result := MakeElementsUnique(startSlice)
	fmt.Println(result)

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
	for {
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

type IDs struct {
	ID int
}

func MakeElementsUnique(slice []IDs) []IDs {
	unique := UniqueValuesSlice(slice)
	sorted := SortSlice(unique)

	return sorted
}

func CreateSlice() []IDs {
	slice := []IDs{}

	for i := 0; i < 20; i++ {
		new_id := IDs{ID: rand.IntN(50)}
		slice = append(slice, new_id)
	}

	fmt.Println(slice)
	return slice
}

func UniqueValuesSlice(slice []IDs) []IDs {
	var uniqueElemsSlice []IDs
	for _, elem := range slice {
		if slices.Contains(uniqueElemsSlice, elem) {
			continue
		} else {
			uniqueElemsSlice = append(uniqueElemsSlice, elem)
		}
	}
	return uniqueElemsSlice
}

func SortSlice(slice []IDs) []IDs {
	len := len(slice)
	for i := 0; i < len-1; i++ {
		swapped := false
		for j := 0; j < len-i-1; j++ {
			if slice[j].ID > slice[j+1].ID {
				slice[j], slice[j+1] = slice[j+1], slice[j]
				swapped = true
			}

		}
		if !swapped {
			break
		}
	}
	return slice
}
