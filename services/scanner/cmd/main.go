package main

import (
	"context"
	"cryptonews/shared/contracts"
	"cryptonews/shared/events"
	"encoding/json"
	//"github.com/k0kubun/pp/v3"
	"log"
	"os"
	"time"

	"github.com/k0kubun/pp/v3"
	"github.com/segmentio/kafka-go"
)

func main() {
	broker := getEnv("KAFKA_BROKERS", "kafka:9092")

	writer := &kafka.Writer{
		Addr:         kafka.TCP(broker),
		Topic:        events.TopicSourceScanRequested,
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireOne,
		Async:        false,
	}
	pp.Println(writer.Topic)
	defer func() {
		if err := writer.Close(); err != nil {
			log.Printf("close writer error: %v", err)
		}
	}()

	log.Printf("scanner started, broker=%s, topic=%s", broker, events.TopicSourceScanRequested)

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		message := contracts.ScanSourceRequested{
			SourceID:    "coindesk",
			SourceType:  "rss",
			URL:         "https://www.coindesk.com/arc/outboundfeeds/rss/",
			RequestedAt: time.Now().UTC().Format(time.RFC3339),
		}

		payload, err := json.Marshal(message)
		if err != nil {
			log.Printf("marshal error: %v", err)
			<-ticker.C
			continue
		}

		err = writer.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(message.SourceID),
			Value: payload,
		})

		if err != nil {
			log.Printf("publish error: %v", err)
		} else {
			log.Printf("published: %s", payload)
		}

		<-ticker.C
	}
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}
