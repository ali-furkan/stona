package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	return port
}

func init() {
	go godotenv.Load(".env")
}
