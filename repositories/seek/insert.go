package seek

import (
	"context"

	"aussie-jobs/seek"

	"gorm.io/gorm/clause"
)

func (sk seekRepo) InsertJob(ctx context.Context, jobs []seek.SummarizedData) error {
	if err := sk.db.Clauses(clause.OnConflict{DoNothing: true}).WithContext(ctx).Table("jobs").Create(&jobs).Error; err != nil {
		return err
	}
	return nil
}
