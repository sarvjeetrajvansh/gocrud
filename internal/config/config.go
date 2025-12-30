package config

import (
	"fmt"
	"os"
)

type Config struct {
	AppName  string
	HTTPPort string

	DBDSN string

	OtelEndpoint  string
	SamplingRatio string

	RabbitMQURL string
}

func Load() *Config {
	return &Config{
		AppName:  os.Getenv("APP_NAME"),
		HTTPPort: os.Getenv("HTTP_PORT"),

		DBDSN: fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=%s",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_SSLMODE"),
		),

		OtelEndpoint:  os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT"),
		SamplingRatio: os.Getenv("OTEL_SAMPLING_RATIO"),

		RabbitMQURL: fmt.Sprintf(
			"amqp://%s:%s@%s:%s/",
			os.Getenv("RABBITMQ_USER"),
			os.Getenv("RABBITMQ_PASSWORD"),
			os.Getenv("RABBITMQ_HOST"),
			os.Getenv("RABBITMQ_PORT"),
		),
	}
}
