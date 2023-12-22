package main

import (
	"fmt"
	"log"

	"aussie-jobs/controllers"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("unable to load dotenv %s", err)
	}

	if err := controllers.InitRouter().Run(); err != nil {
		fmt.Println("unable to init routes", err)
	}
}
