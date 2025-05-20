package model

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	FindCall func(db *gorm.DB)
)

type (
	FindCallFunc func(db *gorm.DB)
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

type dbModel struct {
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

func RawsByCondition(tx *gorm.DB, query string, args []interface{}) (int64, error) {
	var (
		count = RowsCount{}
		field = `count(1) as total`
	)

	if tx == nil {
		tx = EmbedDao.DB()
	}

	db := tx.Model(new(T)).Select(field).Where(query, args...)
	if err := db.Find(&count); err != nil {
		return count.Total, err.Error
	}

	return count.Total, nil
}

func CreateOne[T any](tx *gorm.DB, obj *T) (err error) {
	if err := tx.Model(new(T)).Create(obj); err.Error != nil {
		return err.Error
	}
	return
}

func CreateOneInfo[T any](tx *gorm.DB, obj *T) (err error) {
	if err := tx.Model(new(T)).Create(obj); err.Error != nil {
		return err.Error
	}
	return
}

func CreateOneQuery[T any](tx *gorm.DB, obj *T) (err error) {
	if err := tx.Model(new(T)).Create(obj); err.Error != nil {
		return err.Error
	}
	return
}

func DeleteAll[T any](tx *gorm.DB, query string, args []interface{}) error {
	return tx.Where(query, args...).Delete(new(T)).Error
}

func Delete[T any](tx *gorm.DB, query string, args []interface{}) error {
	return tx.Where(query, args...).Delete(new(T)).Error
}

func DeleteQuery[T any](tx *gorm.DB, query string, args []interface{}) error {
	return tx.Where(query, args...).Delete(new(T)).Error
}

func DeleteBAll[T any](tx *gorm.DB, query string, args []interface{}) error {
	return tx.Where(query, args...).Delete(new(T)).Error
}

func DeleteQuerys[T any](tx *gorm.DB, query string, args []interface{}) error {
	return tx.Where(query, args...).Delete(new(T)).Error
}

func CreateOnes[T any](tx *gorm.DB, obj *T) (err error) {
	if err := tx.Model(new(T)).Create(obj); err.Error != nil {
		return err.Error
	}
	return
}

func CreateBatch[T any](tx *gorm.DB, obj *T) (err error) {
	if err := tx.Model(new(T)).Create(obj); err.Error != nil {
		return err.Error
	}
	return
}
