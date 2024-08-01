package main

import (
	"context"
	"fmt"
	"hw8/enteties"
	"hw8/generator"
	"sync"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	roundsChan := make(chan enteties.Round)
	answersChan := make(chan enteties.PlayerAnswer)

	go generator.GenerateRounds(ctx, roundsChan)

	players := enteties.GeneratePlayers()

	var wg sync.WaitGroup
	for _, player := range players {
		wg.Add(1)
		go func(p enteties.Player) {
			defer wg.Done()
			enteties.PlayerGame(ctx, p.ID, p.Name, roundsChan, answersChan)
		}(player)
	}

	go func() {
		wg.Wait()
		close(answersChan)
	}()

	// Simulate main game loop
	results := make(map[int]int)
	for answer := range answersChan {
		results[answer.PlayerID] += 1
		fmt.Printf("Player %d answered: %d\n", answer.PlayerID, answer.Answer)
	}

	// Print final results
	for id, score := range results {
		fmt.Printf("Player %d scored: %d\n", id, score)
	}
}
