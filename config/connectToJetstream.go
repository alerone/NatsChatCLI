package config

import (
	"github.com/nats-io/nats.go/jetstream"
)

var Js jetstream.JetStream

func CreateJetStream() error {
	js, err := jetstream.New(NatsConnection)
	if err != nil {
		return err
	}
	Js = js
	return nil
}
