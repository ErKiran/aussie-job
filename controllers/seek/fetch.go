package seek

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"

	"aussie-jobs/seek"

	"github.com/gin-gonic/gin"
)

func (sc seekController) SearchJob(ctx *gin.Context) {
	title := ctx.Query("titles")

	titles := strings.Split(title, ",")

	var jobs []seek.SummarizedData
	var wg sync.WaitGroup
	var mu sync.Mutex // Mutex to protect access to the jobs slice

	for _, title := range titles {
		wg.Add(1)
		go func(title string) {
			defer wg.Done()
			job, err := sc.seek.SearchJobs(ctx, title)
			if err != nil {
				fmt.Println("error", err)
				return
			}

			// Use a mutex to safely append to the jobs slice
			mu.Lock()
			jobs = append(jobs, job...)
			mu.Unlock()
		}(title)
	}

	wg.Wait()

	if err := sc.jobRepo.InsertJob(ctx, jobs); err != nil {
		fmt.Println("fucking error on insert", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	jobsJs, _ := json.MarshalIndent(jobs, "", " ")

	if err := os.WriteFile("output.json", jobsJs, 0o644); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("total", len(jobs))
}
