package main

import (
	"fmt"
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
	fmt.Printf("Result of finding is:%v", result)
}

func main() {
	textEditor := NewTextEditor()
	textEditor.AddLine("London is the capital of Great Britain.")
	textEditor.AddLine("London is a huge city.")
	textEditor.AddLine("London city has a river.")

	textEditor.PrintTextEditorInfo()
	textEditor.SearchLinesByWords("London")
}
