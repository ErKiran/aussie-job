package controllers

import (
	"aussie-jobs/controllers/seek"

	"gorm.io/gorm"
)

type controller struct {
	SeekController seek.SeekController
}

func NewController(db *gorm.DB) *controller {
	return &controller{
		SeekController: seek.NewSeekController(db),
	}
}
