package main

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"sort"
	"strings"
)

func main() {
	fmt.Println("*******************FIRST****************************")
	fmt.Println("Please input string to find:")
	var userInputToFindSlice string
	fmt.Scan(&userInputToFindSlice)
	fmt.Printf("Your word to find: %s.\n", userInputToFindSlice)
	createdTextByUser := CreateText()
	FindSliceInText(createdTextByUser, userInputToFindSlice)

	fmt.Println("*******************SECOND****************************")
	startSlice := CreateSlice()
	result := MakeElementsUnique(startSlice)
	fmt.Println(result)

}

func FindSliceInText(Slice []string, wordToFind string) {
	fmt.Println("******************************")
	fmt.Println("Results:")
	found := false
	wordToFind = strings.ToLower(wordToFind)
	for i, val := range Slice {
		if strings.Contains(strings.ToLower(val), wordToFind) {
			fmt.Printf("Line %d: %s\n", i+1, val)
			found = true
		}
	}
	if !found {
		fmt.Println("No matches found.")
	}
}

func CreateText() []string {
	var userCreatedText []string
	var userInput string
	fmt.Println("******************************")
	fmt.Println("Input words, for new word press Enter. To stop press 'q'")
	for {
		fmt.Print(">")
		fmt.Scan(&userInput)
		if userInput == "q" {
			break
		}
		userCreatedText = append(userCreatedText, userInput)
	}
	fmt.Println("******************************")
	fmt.Println(userCreatedText)
	return userCreatedText
}

type IDs struct {
	ID int
}

func MakeElementsUnique(slice []IDs) []IDs {
	unique := UniqueValuesSlice(slice)
	sort.Slice(unique, func(i, j int) bool { return unique[i].ID < unique[j].ID })

	return unique
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

// func SortSlice(slice []IDs) []IDs {
// 	len := len(slice)
// 	for i := 0; i < len-1; i++ {
// 		swapped := false
// 		for j := 0; j < len-i-1; j++ {
// 			if slice[j].ID > slice[j+1].ID {
// 				slice[j], slice[j+1] = slice[j+1], slice[j]
// 				swapped = true
// 			}

// 		}
// 		if !swapped {
// 			break
// 		}
// 	}
// 	return slice
// }
