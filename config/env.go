package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error: ", err)
	}
}

func GetEnv(envName string) string {
	return os.Getenv(envName)
}
