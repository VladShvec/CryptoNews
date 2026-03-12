package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"cryptonews/shared/contracts"
	"cryptonews/shared/events"

	"github.com/segmentio/kafka-go"
)

func main() {
	broker := getEnv("KAFKA_BROKERS", "kafka:9092")
	groupID := getEnv("KAFKA_GROUP_ID", "parser-group")

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{broker},
		GroupID:  groupID,
		Topic:    events.TopicSourceScanRequested,
		MinBytes: 1,
		MaxBytes: 10e6,
	})

	writer := &kafka.Writer{
		Addr:         kafka.TCP(broker),
		Topic:        events.TopicArticleParsed,
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireOne,
		Async:        false,
	}

	defer func() {
		if err := reader.Close(); err != nil {
			log.Printf("reader close error: %v", err)
		}
		if err := writer.Close(); err != nil {
			log.Printf("writer close error: %v", err)
		}
	}()

	log.Printf(
		"parser started, broker=%s, consume_topic=%s, produce_topic=%s",
		broker,
		events.TopicSourceScanRequested,
		events.TopicArticleParsed,
	)

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("read error: %v", err)
			time.Sleep(2 * time.Second)
			continue
		}

		var scanEvent contracts.ScanSourceRequested
		if err := json.Unmarshal(msg.Value, &scanEvent); err != nil {
			log.Printf("unmarshal error: %v", err)
			continue
		}

		log.Printf("received scan event: %+v", scanEvent)

		parsedEvent := contracts.ArticleParsed{
			SourceID:    scanEvent.SourceID,
			Title:       "Test article from parser",
			URL:         scanEvent.URL + "/article-1",
			PublishedAt: time.Now().UTC().Format(time.RFC3339),
			Content:     "This is a stub parsed article payload",
		}

		payload, err := json.Marshal(parsedEvent)
		if err != nil {
			log.Printf("marshal parsed event error: %v", err)
			continue
		}

		err = writer.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(parsedEvent.SourceID),
			Value: payload,
		})
		if err != nil {
			log.Printf("publish parsed event error: %v", err)
			continue
		}

		log.Printf("published parsed event: %s", payload)
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
