package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type TextEditor struct {
	lines []string
	index map[string][]int
}

func (t TextEditor) PrintTextEditorInfo() {
	fmt.Printf("TEXT LINES: %v\nWORDS INDEXES: %v\n", t.lines, t.index)
}

func NewTextEditor() *TextEditor {
	return &TextEditor{
		index: make(map[string][]int),
	}
}

// For reading from .txt file
func (t *TextEditor) ReadFileByLine(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	line, err := Readln(reader)
	for err == nil {
		t.lines = append(t.lines, line)
		t.IndexLine(len(t.lines) - 1)
		line, err = Readln(reader)
	}
}

func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

// For simple adding
func (t *TextEditor) AddLine(line string) {
	t.lines = append(t.lines, line)
	t.IndexLine(len(t.lines) - 1)
}

func (t *TextEditor) IndexLine(lineIndex int) {
	words := strings.Fields(t.lines[lineIndex])
	for _, word := range words {
		word = strings.ToLower(word)
		t.index[word] = append(t.index[word], lineIndex)
	}
}

func (t TextEditor) SearchLinesByWords(word string) {
	result := []string{}
	wordLower := strings.ToLower(word)
	for _, lineIndex := range t.index[wordLower] {
		result = append(result, t.lines[lineIndex])
	}
	fmt.Printf("Result of finding is: %v", result)
}

func main() {
	textEditor := NewTextEditor()
	path := "E:\\Go programming\\golang_projector_homeworks\\hw_5\\task.txt"

	textEditor.ReadFileByLine(path)
	var userInput string
	for {
		fmt.Print("\nInput word you want to find.To exit write quit: ")
		fmt.Scan(&userInput)
		if userInput != "quit" {
			textEditor.SearchLinesByWords(userInput)
		} else {
			break
		}

	}
}
