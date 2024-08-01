package generator

import (
	"context"
	"hw8/enteties"
	"math/rand"
	"time"
)

func GenerateRounds(ctx context.Context, roundsChan chan<- enteties.Round) {
	questions := []enteties.Question{
		{
			Text:    "What is the capital of France?",
			Answers: map[int]string{1: "Paris", 2: "London", 3: "Berlin", 4: "Madrid"},
			Correct: 0,
		},
		{
			Text:    "Which planet is known as the Red Planet?",
			Answers: map[int]string{1: "Earth", 2: "Mars", 3: "Jupiter", 4: "Saturn"},
			Correct: 1,
		},
		{
			Text:    "Who wrote 'To Kill a Mockingbird'?",
			Answers: map[int]string{1: "Harper Lee", 2: "Mark Twain", 3: "F. Scott Fitzgerald", 4: "Ernest Hemingway"},
			Correct: 0,
		},
		{
			Text:    "What is the largest ocean on Earth?",
			Answers: map[int]string{1: "Atlantic Ocean", 2: "Indian Ocean", 3: "Arctic Ocean", 4: "Pacific Ocean"},
			Correct: 3,
		},
		{
			Text:    "What is the chemical symbol for water?",
			Answers: map[int]string{1: "O2", 2: "H2O", 3: "CO2", 4: "HO"},
			Correct: 1,
		},
		{
			Text:    "Who was the first President of the United States?",
			Answers: map[int]string{1: "George Washington", 2: "Thomas Jefferson", 3: "Abraham Lincoln", 4: "John Adams"},
			Correct: 0,
		},
		{
			Text:    "What is the speed of light?",
			Answers: map[int]string{1: "300,000 km/s", 2: "150,000 km/s", 3: "100,000 km/s", 4: "200,000 km/s"},
			Correct: 0,
		},
		{
			Text:    "Which language is used to create web pages?",
			Answers: map[int]string{1: "Python", 2: "Java", 3: "HTML", 4: "C++"},
			Correct: 2,
		},
		{
			Text:    "What is the hardest natural substance on Earth?",
			Answers: map[int]string{1: "Gold", 2: "Iron", 3: "Diamond", 4: "Silver"},
			Correct: 2,
		},
	}

	for {
		select {
		case <-ctx.Done():
			close(roundsChan)
			return
		case <-time.After(10 * time.Second):
			question := questions[rand.Intn(len(questions))]
			round := enteties.Round{Question: question}
			roundsChan <- round
		}
	}
}
