package seek

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

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

	// titles := []string{"Golang", "React", "Node", "Backend", "Frontend", "Java", "Javascript", "Kubernetes", "PHP", "Python", "RUST", "SQL", "TYPESCRIPT", "DEVOPS", "DATASIENCE", "CYBERSECURITY"}

	titles := []string{"Golang"}

	var jobs []seek.SummarizedData
	for _, title := range titles {
		job, err := seeks.SearchJobs(ctx, title)
		if err != nil {
			fmt.Println("error", err)
		}
		jobs = append(jobs, job...)
	}

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
