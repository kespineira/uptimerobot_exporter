package config

import (
	"log"
	"os"
)

var (
	ApiKey string
	Port   string
)

func LoadConfig() {
	ApiKey = os.Getenv("UPTIMEROBOT_API_KEY")
	if ApiKey == "" {
		log.Fatal("UPTIMEROBOT_API_KEY not set")
	}
	Port = os.Getenv("PORT")
	if Port == "" {
		Port = "8080"
	}
}
