package app

import (
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Port           string `env:"PORT" env-default:"80"`
	DatabaseSource string `env:"DB_SOURCE"`
	LogLevel       string `env:"LOG_LEVEL" env-default:"debug"`
	YandexToken    string `env:"YANDEX_TOKEN"`
}

func configure() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(".env file was not found. Using env vars")
	} else {
		log.Println(".env file was successfuly found and parsed")
	}
	var config Config
	err = cleanenv.ReadEnv(&config)
	if err != nil {
		log.Fatal(fmt.Errorf("parse env vars error: %w", err))
	}
	return &config
}
