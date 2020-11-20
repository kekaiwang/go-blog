package model

type ArticleRelation struct {
	Id        int64 `gorm:"primary_key"`
	ArticleId int64 `gorm:"article_id"` // 文章ID
	TagId     int64 `gorm:"tag_id"`     // 标签ID
}

var articleRelation *ArticleRelation

func ArticleRelationModel() *ArticleRelation {
	return articleRelation
}

func (ar *ArticleRelation) TableName() string {
	return `article_relation`
}
