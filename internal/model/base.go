package model

import "gorm.io/gorm"

type (
	FindCall func(db *gorm.DB)
)

type RowsCount struct {
	Total int64 `gorm:"column:total"`
}

const (
	MaxBatchSize = 500
)
