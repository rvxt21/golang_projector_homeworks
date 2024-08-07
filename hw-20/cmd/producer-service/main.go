package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
)

const orangesTopic = "oranges"

func main() {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", orangesTopic, 0)
	if err != nil {
		log.Fatal().Msgf("failed to dial leader: %s", err)
	}
	defer conn.Close()
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	idService := IDGeneratorService{}
	for {
		select {
		case <-ticker.C:
			id := idService.GenerateID()
			size := rand.Intn(300)
			randomMessage := fmt.Sprintf(`{"OrangeID": %d, "Size": %d}`, id, size)
			_, err = conn.WriteMessages(
				kafka.Message{Value: []byte(randomMessage)},
			)

			if err != nil {
				log.Fatal().Err(err).Msg("failed to write kafka messages")
			}
		}
	}
}
