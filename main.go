package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"aussie-jobs/seek"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("unable to load dotenv %s", err)
	}

	seek := seek.NewSeek()

	jobs, err := seek.SearchJobs(context.TODO(), "Flutter")
	if err != nil {
		fmt.Println("error", err)
	}

	jobsJs, _ := json.MarshalIndent(jobs, "", " ")

	fmt.Println("jobs...", string(jobsJs))

	if err = os.WriteFile("output.json", jobsJs, 0o644); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("total", len(jobs))
}
