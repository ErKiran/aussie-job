package seek

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"

	repo "aussie-jobs/repositories/seek"
	"aussie-jobs/seek"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type seekController struct {
	seek    seek.Seek
	jobRepo repo.SeekRepo
}

type SeekController interface {
	SearchJob(ctx *gin.Context)
}

func NewSeekController(db *gorm.DB) seekController {
	return seekController{
		seek:    seek.NewSeek(),
		jobRepo: repo.NewSeekRepo(db),
	}
}

func (sc seekController) SearchJob(ctx *gin.Context) {
	seeks := seek.NewSeek()

	title := ctx.Query("titles")

	titles := strings.Split(title, ",")

	var jobs []seek.SummarizedData
	var wg sync.WaitGroup
	var mu sync.Mutex // Mutex to protect access to the jobs slice

	for _, title := range titles {
		wg.Add(1)
		go func(title string) {
			defer wg.Done()
			job, err := seeks.SearchJobs(ctx, title)
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
