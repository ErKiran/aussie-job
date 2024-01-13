package seek

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"aussie-jobs/seek"

	"github.com/gin-gonic/gin"
)

func (sc seekController) InitSearch(ctx context.Context, titles []string) error {
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
		return err
	}

	fmt.Println("total", len(jobs))
	return nil
}

func (sc seekController) SearchJob(ctx *gin.Context) {
	title := ctx.Query("titles")

	titles := strings.Split(title, ",")

	if err := sc.InitSearch(ctx, titles); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
}
