package tag

import (
	"time"

	"github.com/kekaiwang/go-blog/internal/model"
)

type GetTagListRequest struct {
	Limit  int64  `json:"limit"`
	Page   int64  `json:"page"`
	Offset int64  `json:"offset"`
	Name   string `json:"name"`
}

type GetTagListResponse struct {
	Data  []*model.Tag `json:"data"`
	Total int64        `json:"total"`
}

type UpdateTagRequest struct {
	ID         int64     `gorm:"id" json:"id"`
	Name       string    `gorm:"name" json:"name"`               //名称
	RouterLink string    `gorm:"router_link" json:"router_link"` // 路由链接
	UseTimes   int64     `gorm:"use_times" json:"use_times"`     // tag链接文章数量
	Status     int       `gorm:"status" json:"status"`           // 1:正常 2:禁用
	Created    time.Time `gorm:"created" json:"created"`         // created
	Updated    time.Time `gorm:"updated" json:"updated"`
}

type CreateTagRequest struct {
	Name       string    `gorm:"name" json:"name"`               //名称
	RouterLink string    `gorm:"router_link" json:"router_link"` // 路由链接
	UseTimes   int64     `gorm:"use_times" json:"use_times"`     // tag链接文章数量
	Status     int       `gorm:"status" json:"status"`           // 1:正常 2:禁用
	Created    time.Time `gorm:"created" json:"created"`
	Updated    time.Time `gorm:"updated" json:"updated"`
}
