package seek

import (
	"context"

	"aussie-jobs/seek"
)

func (sk seekRepo) InsertJob(ctx context.Context, jobs []seek.SummarizedData) error {
	if err := sk.db.WithContext(ctx).Table("jobs").Create(&jobs).Error; err != nil {
		return err
	}
	return nil
}
