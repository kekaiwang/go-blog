package model

import (
	"time"

	"github.com/kekaiwang/go-blog/pkg/drives"
)

type PageInfo struct {
	Id        int64     `gorm:"primary_key"`
	Name      string    `gorm:"name" json:"name"`             // 页面名字
	Slug      string    `gorm:"slug" json:"slug"`             // 页面标识,路由地址
	Content   string    `gorm:"content" json:"content"`       // 内容
	Markdown  string    `gorm:"markdown" json:"markdown"`     // markdown info
	Status    int       `gorm:"status" json:"status"`         //1:正常 2:禁用 3:已删除
	Preview   int64     `gorm:"preview" json:"preview"`       // 阅读数量
	CreatedAt time.Time `gorm:"created_at" json:"created_at"` // created time
	UpdatedAt time.Time `gorm:"updated_at" json:"updated_at"` // updated time
}

var pageInfo *PageInfo

func PageInfoModel() *PageInfo {
	return pageInfo
}

// TableName. table name page info
func (pi *PageInfo) TableName() string {
	return `page_info`
}

// TableName. table name page info
func (pi *PageInfo) TableNames() string {
	return `page_info`
}

// Create insert pageInfo
func (pi *PageInfo) Create(pageInfo *PageInfo) error {
	return drives.BlogDB.Create(pageInfo).Error
}

// GetPageBySlug get page by slug
func (pi *PageInfo) GetPageBySlug(slug string) (*PageInfo, error) {
	page := &PageInfo{}

	err := drives.BlogDB.Table(pi.TableName()).Where("slug = ? ", slug).First(&page).Error
	if err != nil {
		return nil, err
	}

	return page, nil
}

// Create insert pageInfo
func (pi *PageInfo) CreatePa(pageInfo *PageInfo) error {
	return drives.BlogDB.Create(pageInfo).Error
}

// Create insert pageInfo
func (pi *PageInfo) CreatePage(pageInfo *PageInfo) error {
	return drives.BlogDB.Create(pageInfo).Error
}

// Create insert pageInfo
func (pi *PageInfo) CreateIn(pageInfo *PageInfo) error {
	return drives.BlogDB.Create(pageInfo).Error
}
