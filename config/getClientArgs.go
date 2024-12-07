package config

import (
	"errors"
	"natsChat/models"
	"os"
)

var ClientConn models.ClientConnection

func GetClientArgs() error {
	if len(os.Args) < 4 {
		return errors.New("os len error")
	}

	server := os.Args[1]
	channel := os.Args[2]
	name := os.Args[3]

	ClientConn = models.ClientConnection{
		Server:  server,
		Channel: channel,
		Name:    name,
	}
	return nil
}
