package main

import (
	"fmt"
	"log"
	"os"

	"aussie-jobs/repositories"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("unable to load dotenv %s", err)
	}
	fmt.Println("port", os.Getenv("PORT"))
	if err := repositories.Migrate(); err != nil {
		log.Fatalf("Error while running migrations, %v", err)
	}
}
