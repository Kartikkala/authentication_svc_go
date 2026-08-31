package config

import (
	"os"
)

func getEnvOrDefault(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok && val != "" {
		return val
	}
	return fallback
}

type DatabaseConfig struct {
	User     string
	Password string
	Hostname string
	DBName   string
	Port     uint16
	Timezone string
}

type ApplicationConfig struct {
	RESTPort    uint16
	HostAddress string
}

type EmailConfig struct {
	Email    string
	Password string
	Port     uint16
}

type JWTConfig struct {
	Secret     string
	ExpiryHour int
}


type NATSConfig struct {
	URL string
}

type Config struct {
	Database DatabaseConfig
	App      ApplicationConfig
	SMTP     EmailConfig
	JWT      JWTConfig
	NATS     NATSConfig
}

func NewConfig() *Config {

	return &Config{
		Database: DatabaseConfig{
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASS"),
			DBName:   os.Getenv("POSTGRES_DBNAME"),
			Hostname: getEnvOrDefault("POSTGRES_HOSTNAME", "127.0.0.1"),
			Port:     uint16(5432),
			Timezone: getEnvOrDefault("POSTGRES_TIMEZONE", "Asia/Kolkata"),
		},
		App: ApplicationConfig{
			RESTPort:    8080,
			HostAddress: getEnvOrDefault("HOST_ADDRESS", "127.0.0.1"),
		},
		SMTP: EmailConfig{
			Email:    os.Getenv("EMAIL_ADDRESS"),
			Password: os.Getenv("EMAIL_PASSWORD"),
			Port:     uint16(587),
		},
		JWT: JWTConfig{
			Secret:     os.Getenv("JWT_SECRET"),
			ExpiryHour: 24,
		},
		NATS: NATSConfig{
			URL: getEnvOrDefault("NATS_URL", "nats://127.0.0.1:4222"),
		},
	}
}
