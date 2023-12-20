package seek

import (
	"context"
	"os"

	"aussie-jobs/utils"
)

const (
	SEARCH = "search"
)

type Seek interface {
	SearchJobs(ctx context.Context, keyword string) ([]SummarizedData, error)
}

type SeekAPI struct {
	client *utils.Client
}

func NewSeek() Seek {
	client := utils.NewClient(nil, os.Getenv("SEEK"), "")

	seek := &SeekAPI{
		client: client,
	}
	return seek
}
