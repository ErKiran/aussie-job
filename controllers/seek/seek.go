package seek

import (
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
	AnalyzeJobs(ctx *gin.Context)
}

func NewSeekController(db *gorm.DB) seekController {
	return seekController{
		seek:    seek.NewSeek(),
		jobRepo: repo.NewSeekRepo(db),
	}
}
