/**
 * Author: kekai wang
 */
package model

import "time"

//Article 文章
type Article struct {
	Id          int64     `gorm:"primary_key" json:"id"`
	Author      string    `gorm:"author" json:"author"`             // 作者
	Count       int64     `gorm:"count" json:"count"`               // 评论数量
	Markdown    string    `gorm:"markdown"`                         // 文章内容markdown格式
	Content     string    `gorm:"content" json:"content"`           // 文章内容
	CategoryId  int64     `gorm:"category_id" json:"category_id"`   // 分类
	TagIds      string    `gorm:"tag_ids" json:"tag_ids"`           // 标签分类
	Excerpt     string    `gorm:"excerpt" json:"excerpt"`           // 预览信息
	Previous    string    `gorm:"previous" json:"previous"`         // 前一篇
	Next        string    `gorm:"next" json:"next"`                 // 后一篇
	Preview     int64     `gorm:"preview" json:"preview"`           // 浏览数量
	Thumb       string    `gorm:"thumb" json:"thumb"`               // 缩略图
	Slug        string    `gorm:"slug" json:"slug"`                 // 路由地址
	IsDraft     int       `gorm:"is_draft" json:"is_draft"`         // 1：草稿 2:已发布 3:已删除
	EditTime    time.Time `gorm:"edit_time" json:"edit_time"`       //编辑时间
	DisplayTime time.Time `gorm:"display_time" json:"display_time"` // 显示时间
	CreatedAt   time.Time `gorm:"created_at" json:"created_at"`     //创建时间
	UpdatedAt   time.Time `gorm:"updated_at" json:"updated_at"`     // 更新时间

}

var article *Article

func ArticleModel() *Article {
	return article
}

func (a *Article) TableName() string {
	return `article`
}

func (a *Article) GetArticleBySlug() (*Article, error) {
	article := &Article{}
}
