package config

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

// EnvConfig menyimpan konfigurasi aplikasi
type EnvConfig struct {
	ServerPort string `env:"SERVER_PORT" envDefault:"8080"`
	DBHost     string `env:"DB_HOST,required"`
	DBName     string `env:"DB_NAME,required"`
	DBUser     string `env:"DB_USER,required"`
	DBPassword string `env:"DB_PASSWORD,required"`
	DBSSLMode  string `env:"DB_SSLMODE" envDefault:"disable"`
}

// LoadEnv membaca file .env
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}
}

// NewEnvConfig menginisialisasi dan mem-parsing konfigurasi dari environment
func NewEnvConfig() *EnvConfig {
	LoadEnv() // Load .env file

	config := &EnvConfig{}
	if err := env.Parse(config); err != nil {
		log.Fatalf("Error parsing environment variables: %v", err)
	}

	return config
}
