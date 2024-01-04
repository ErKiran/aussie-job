package seek

import (
	"context"

	"aussie-jobs/seek"

	"gorm.io/gorm"
)

type SeekRepo interface {
	InsertJob(ctx context.Context, jobs []seek.SummarizedData) error
	JobTitle(ctx context.Context) ([]JobLocation, error)
	CompanyTitle(ctx context.Context) ([]Company, error)
}

type seekRepo struct {
	db *gorm.DB
}

func NewSeekRepo(db *gorm.DB) SeekRepo {
	return seekRepo{
		db: db,
	}
}
