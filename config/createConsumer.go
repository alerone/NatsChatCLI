package config

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/nats-io/nats.go/jetstream"
)

var Consumer jetstream.Consumer

func CreateConsumer() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	channelName := strings.ReplaceAll(ClientConn.Channel, ".", "_")

	consumerName := fmt.Sprintf("%s_%s", channelName, ClientConn.Name)

	_, err := ChatStream.Consumer(ctx, consumerName)
	if err != nil {
		if err == jetstream.ErrConsumerNotFound {
			consumer, err := ChatStream.CreateConsumer(ctx, jetstream.ConsumerConfig{
				Name:          consumerName,
				Durable:       consumerName,
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
	}
	return fmt.Errorf("user %s already exists in channel %s", ClientConn.Name, ClientConn.Channel)
}

func DeleteConsumer() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	ChatStream.DeleteConsumer(ctx, Consumer.CachedInfo().Name)
}
