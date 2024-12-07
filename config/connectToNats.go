package config

import (
	"github.com/nats-io/nats.go"
)

var NatsConnection *nats.Conn

func ConnectToNats() error {
	conn, err := nats.Connect(ClientConn.Server)

	if err != nil {
		return err
	}
	NatsConnection = conn
	return nil

}
