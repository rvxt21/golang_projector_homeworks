package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/segmentio/kafka-go"
)

const orangesTopic = "oranges"

func main() {
	ctx := context.Background()

	storage := NewInMemStorage()
	analytics := NewOrangesAnalytics()
	service := &OrangeService{
		Storage:   storage,
		Analytics: *analytics,
	}

	go func() {
		http.ListenAndServe(":8080", OrangesAnalyticsHandler(service))
	}()

	go func() {
		http.ListenAndServe(":8081", AllOrangesHandler(service))
	}()
	ticker := time.NewTicker(time.Second * 10)
	url := "http://localhost:8080"
	defer ticker.Stop()

	kafkaReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   orangesTopic,
	})

	for {
		select {
		case <-ticker.C:
			CallHandler(url)
		default:

			msg, err := kafkaReader.ReadMessage(ctx)
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to read message")
			}
			log.Info().Str("msg", string(msg.Value)).Msg("Got message from kafka")

			var oe OrangeEvent

			if err := json.Unmarshal(msg.Value, &oe); err != nil {
				log.Warn().Err(err).Msg("Failed to decode message")
				continue
			}

			if err := service.ConsumeOrangeEvent(oe); err != nil {
				log.Warn().Err(err).Msg("Failed to consume")
			}
		}
	}

}
