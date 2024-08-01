package enteties

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
)

type PlayerAnswer struct {
	PlayerID int
	Answer   int
	Scores   int
}

type Player struct {
	ID     int
	Name   string
	Scores int
}

var (
	lastPlayerID int
	mu           sync.Mutex
)

func NewPlayer(name string) *Player {
	mu.Lock()
	defer mu.Unlock()
	lastPlayerID++
	return &Player{
		ID:     lastPlayerID,
		Name:   name,
		Scores: 0,
	}
}

func (p Player) PrintPlayerInfo() {
	fmt.Printf("Player name: %v, player ID: %d, players scores: %d.\n", p.Name, p.ID, p.Scores)
}

func GeneratePlayers() []Player {
	players := []Player{}
	player1 := NewPlayer("Anastasiia")
	player1.PrintPlayerInfo()
	player2 := NewPlayer("Ivan")
	player2.PrintPlayerInfo()
	player3 := NewPlayer("Sofiia")
	player3.PrintPlayerInfo()
	players = append(players, *player1)
	players = append(players, *player2)
	players = append(players, *player3)
	return players
}

func PlayerGame(ctx context.Context, id int, name string, roundsChan <-chan Round, answersChan chan<- PlayerAnswer) {
	for {
		select {
		case <-ctx.Done():
			return
		case round, ok := <-roundsChan:
			if !ok {
				return
			}
			// Simulate answering a question
			answer := rand.Intn(len(round.Question.Answers))
			answersChan <- PlayerAnswer{PlayerID: id, Answer: answer}
		}
	}
}
