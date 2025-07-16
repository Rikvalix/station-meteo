package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	mongoURI := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s/%s?authSource=admin",
		os.Getenv("MONGO_DB_USER"),
		os.Getenv("MONGO_DB_PASSWORD"),
		os.Getenv("MONGO_DB_HOST"),
		os.Getenv("MONGO_DB_PORT"),
		os.Getenv("MONGO_DB_NAME"),
	)
	return mongoURI
}
