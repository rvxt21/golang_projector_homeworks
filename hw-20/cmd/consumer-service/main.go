package main

import (
	"context"
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

	conn.SetReadDeadline(time.Now().Add(time.Second * 30))
	batch := conn.ReadBatch(10e3, 1e6)
	defer batch.Close()

	for {
		msg, err := batch.ReadMessage()
		if err != nil {
			break
		}
		log.Info().Str("msg", string(msg.Value)).Msg("Got message from kafka")
	}
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read batch message")
	}

}
