package main

import (
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	HOST     string `json:"host"`
	PORT     string `json:"port"`
	USER     string `json:"user"`
	DATABASE string `json:"database"`
	PASSWORD string `json:"password"`
}

func LoadConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("unable to load dotenv %s", err)
	}

	return Config{}
}
