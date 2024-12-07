package service

import (
	"fmt"
	"natsChat/config"

	"github.com/nats-io/nats.go"
)

var subscription *nats.Subscription

func Subscribe() error {
	sub, err := config.NatsConnection.Subscribe(config.ClientConn.Channel, func(msg *nats.Msg) {
		fmt.Printf("[%s]: %s\n", msg.Header.Get("name"), string(msg.Data))
	})
	if err != nil {
		return err
	}

	subscription = sub

	return nil
}
