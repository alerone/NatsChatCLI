package config

import (
	"context"
	"fmt"
	"time"

	"github.com/nats-io/nats.go/jetstream"
)

var Consumer jetstream.Consumer

func CreateConsumer() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	consumer, err := ChatStream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		Name:          ClientConn.Name,
		Durable:       ClientConn.Name,
		AckPolicy:     jetstream.AckExplicitPolicy,
		FilterSubject: fmt.Sprintf("chats.%s", ClientConn.Channel),
		DeliverPolicy: jetstream.DeliverAllPolicy,
	})
	if err != nil {
		return err
	}

	Consumer = consumer

	return nil
}

func DeleteConsumer() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	ChatStream.DeleteConsumer(ctx, Consumer.CachedInfo().Name)
}
