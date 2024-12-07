package main

import (
	"bufio"
	"context"
	"log"
	"natsChat/config"
	"natsChat/service"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	err := config.GetClientArgs()
	if err != nil {
		log.Fatal("Usage: chat-cli <NATS_SERVER> <CHANNEL> <YOUR_NAME>")
	}
	err = config.ConnectToNats()

	if err != nil {
		log.Fatalf("Error connecting to NATS server: %v", err)
	}
	err = config.CreateJetStream()
	if err != nil {
		log.Fatalf("Error creating JetStream: %v", err)
	}

	err = config.InitializeStream()
	if err != nil {
		log.Fatalf("Error creating chat stream: %v", err)
	}

	err = config.CreateConsumer()
	if err != nil {
		log.Fatalf("Error creating to Jetstream consumer: %v", err)
	}
}

func main() {

	defer config.DrainNatsConnection()
	defer config.DeleteConsumer()

	_, cancel := context.WithCancel(context.Background())
	log.Printf("Connected to NATS on server %s on channel %s", config.ClientConn.Server, config.ClientConn.Channel)

	service.Consume()
	defer service.ConsumeContextStop()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}

		err := service.PublishText(text)
		if err != nil {
			log.Printf("Error publishing message: %v", err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading input: %v", err)
	}

	// Handle OS signals for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		cancel()
	}()
}
