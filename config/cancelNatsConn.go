package config

func DrainNatsConnection() {
	NatsConnection.Drain()
}
