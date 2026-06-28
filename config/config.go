package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName string
	Version string
	Port    string
}

func LoadConfig() Config {

	err := godotenv.Load()

	if err != nil {

		log.Println("No existe archivo .env, usando configuración por defecto")

	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8081"
	}

	return Config{

		AppName: os.Getenv("APP_NAME"),
		Version: os.Getenv("APP_VERSION"),
		Port:    port,
	}

}
