package model

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	FindCall func(db *gorm.DB)
)

type RowsCount struct {
	Total int64 `gorm:"column:total"`
}

const (
	MaxBatchSize = 500
)

type txModel struct {
	Dao
}

func FindByQueryCondition[T any](field string, query string, args []interface{}, opts ...ExtraOption) ([]*T, error) {
	var results []*T

	db := EmbedDao.DBExtra(EmbedDao.DB(), opts...).Model(new(T)).Select(field).Where(query, args...)

	if err := db.Find(&results); err.Error != nil {
		return results, err.Error
	}

	return results, nil
}

func FindByQueryCondition[T any](field string, query string, args []interface{}, opts ...ExtraOption) ([]*T, error) {
	var results []*T

	db := EmbedDao.DBExtra(EmbedDao.DB(), opts...).Model(new(T)).Select(field).Where(query, args...)

	if err := db.Find(&results); err.Error != nil {
		return results, err.Error
	}

	return results, nil
}

func FindByLockCondition[T any](tx *gorm.DB, field string, query string, args []interface{}) (results []*T, err error) {

	db := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Model(new(T)).Select(field).Where(query, args...)

	if err := db.Find(&results); err.Error != nil {
		return results, err.Error
	}

	return
}

func FindByLockConditionQuery[T any](tx *gorm.DB, field string, query string, args []interface{}) (results []*T, err error) {

	db := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Model(new(T)).Select(field).Where(query, args...)

	if err := db.Find(&results); err.Error != nil {
		return results, err.Error
	}

	return
}
