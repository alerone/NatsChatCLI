package service

import (
	"fmt"
	"natsChat/config"

	"github.com/nats-io/nats.go/jetstream"
)

var consumerContext jetstream.ConsumeContext

func Consume() error {
	cctx, err := config.Consumer.Consume(func(msg jetstream.Msg) {
		fmt.Printf("%s\n", string(msg.Data()))
		msg.Ack()
	})
	if err != nil {
		return err
	}
	consumerContext = cctx
	return nil
}

func ConsumeContextStop() {
	consumerContext.Stop()
}
