package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvData(path string) {
	err := godotenv.Load(path)
	if err != nil {
		log.Fatal("Error while uploading file .env")
	}
}
