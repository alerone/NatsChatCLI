package config

import (
	"context"
	"time"

	"github.com/nats-io/nats.go/jetstream"
)

var ChatStream jetstream.Stream

func InitializeStream() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	stream, err := Js.Stream(ctx, "chats")

	if err != nil {
		stream, err = Js.CreateStream(ctx, jetstream.StreamConfig{
			Name:        "chats",
			Description: "Stream for chatting",
			Subjects:    []string{"chats.>"},
			MaxAge:      time.Hour,
		})
		if err != nil {
			return err
		}

	}
	ChatStream = stream
	return nil
}
