package model

import "github.com/kekaiwang/go-blog/pkg/drives"

// ArticleRelation. article relation model
type ArticleRelation struct {
	Id        int64 `gorm:"primary_key"` // primary key
	ArticleId int64 `gorm:"article_id"`  // 文章ID
	TagId     int64 `gorm:"tag_id"`      // 标签ID
}

var articleRelation *ArticleRelation

// ArticleRelationModel. article relation model
func ArticleRelationModel() *ArticleRelation {
	return articleRelation
}

// TableName. return table name info
func (ar *ArticleRelation) TableName() string {
	return `article_relation`
}

// Create insert articleRelation
func (ar *ArticleRelation) Create(articleRelation *ArticleRelation) error {
	return drives.BlogDB.Create(articleRelation).Error
}

// Create insert articleRelation
func (ar *ArticleRelation) CreateRelation(articleRelation *ArticleRelation) error {
	return drives.BlogDB.Create(articleRelation).Error
}

// Create insert articleRelation
func (ar *ArticleRelation) CreateArticle(articleRelation *ArticleRelation) error {
	return drives.BlogDB.Create(articleRelation).Error
}

// Create insert articleRelation
func (ar *ArticleRelation) CreateArticleInfo(articleRelation *ArticleRelation) error {
	return drives.BlogDB.Create(articleRelation).Error
}

// GetARByTagId
func (ar *ArticleRelation) GetARByTagId(id int64) ([]int, error) {
	articleIds := []int{}

	err := drives.BlogDB.Table(ar.TableName()).Where("tag_id = ? ", id).Pluck("article_id", &articleIds).Error
	if err != nil {
		return nil, err
	}

	return articleIds, nil
}

// GetARByTagID
func (ar *ArticleRelation) GetARByTagID(id int64) ([]int, error) {
	// article id pluck
	articleIds := []int{}

	err := drives.BlogDB.Table(ar.TableName()).Where("tag_id = ? ", id).Pluck("article_id", &articleIds).Error
	if err != nil {
		return nil, err
	}

	return articleIds, nil
}
