package category

import (
	"time"

	"github.com/kekaiwang/go-blog/internal/model"
)

type GetCategoryReq struct {
	Link   string `json:"link"`
	Limit  int64  `json:"limit"`
	Offset int64  `json:"offset"`
}

type GetCategoryRes struct {
	Data  []*CategoryArticle `json:"data"`  // data
	Name  string             `json:"name"`  // name
	Total int64              `json:"total"` // total
	Meta  Meta               `json:"meta"`  // meta
}

type CategoryArticle struct {
	Title       string `json:"title"`        // title
	Slug        string `json:"slug"`         // slug
	DisplayTime string `json:"display_time"` // dispaly
}

type Meta struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type GetTagReq struct {
	Link   string `json:"link"`
	Limit  int64  `json:"limit"`
	Offset int64  `json:"offset"`
}

type AdminCategoryListRequest struct {
	Limit  int64  `json:"limit"`
	Page   int64  `json:"page"`
	Offset int64  `json:"offset"`
	Name   string `json:"name"`
}

type AdminCategoryListResponse struct {
	Data  []*model.Category `json:"data"`
	Total int64             `json:"total"`
}

type UpdateCategoryRequest struct {
	ID         int64     `gorm:"id" json:"id"`
	Name       string    `gorm:"name" json:"name"`               //名称
	RouterLink string    `gorm:"router_link" json:"router_link"` // 路由链接
	Status     int       `gorm:"status" json:"status"`           // 1:正常 2:禁用 3:已删除
	Created    time.Time `gorm:"created_at" json:"created"`
	Updated    time.Time `gorm:"updated_at" json:"updated"`
}

type CreateCategoryRequest struct {
	Name        string    `gorm:"name" json:"name"`                 //名称
	RouterLink  string    `gorm:"router_link" json:"router_link"`   // 路由链接
	LinkArticle int64     `gorm:"link_article" json:"link_article"` // 链接文章数量
	Status      int       `gorm:"status" json:"status"`             // 1:正常 2:禁用 3:已删除
	Created     time.Time `gorm:"created_at" json:"created"`
	Updated     time.Time `gorm:"updated_at" json:"updated"`
}
