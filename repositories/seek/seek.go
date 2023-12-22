package seek

import (
	"context"

	"aussie-jobs/seek"

	"gorm.io/gorm"
)

type SeekRepo interface {
	InsertJob(ctx context.Context, jobs []seek.SummarizedData) error
}

type seekRepo struct {
	db *gorm.DB
}

func NewSeekRepo(db *gorm.DB) SeekRepo {
	return seekRepo{
		db: db,
	}
}
