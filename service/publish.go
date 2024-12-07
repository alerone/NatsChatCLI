package service

import (
	"context"
	"fmt"
	"natsChat/config"
	"time"
)

func PublishText(text string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	msg := []byte(fmt.Sprintf("[%v]: %s", config.ClientConn.Name, text))

	if _, err := config.Js.Publish(ctx, fmt.Sprintf("chats.%s", config.ClientConn.Channel), msg); err != nil {
		return err
	}
	return nil
}
