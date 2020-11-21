package model

import (
	"time"

	"github.com/kekaiwang/go-blog/pkg/drives"
)

type Category struct {
	Id          int64     `gorm:"primary_key"`
	Name        string    `gorm:"name" json:"name"`                 //名称
	RouterLink  string    `gorm:"router_link" json:"router_link"`   // 路由链接
	LinkArticle int64     `gorm:"link_article" json:"link_article"` // 链接文章数量
	Status      int       `gorm:"status" json:"status"`             // 1:正常 2:禁用 3:已删除
	CreatedAt   time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"updated_at" json:"updated_at"`
}

var category *Category

func CategoryModel() *Category {
	return category
}

func (c *Category) TableName() string {
	return `category`
}

//Create insert category
func (c *Category) Create(category *Category) error {
	return drives.BlogDB.Create(category).Error
}

//GetAll
func (c *Category) GetAll() ([]*Category, error) {
	categories := []*Category{}

	err := drives.BlogDB.Table(c.TableName()).Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}

//GetCategoryByRouterLink
func (c *Category) GetCategoryByRouterLink(routerLink string) (*Category, error) {
	category := Category{}

	err := drives.BlogDB.Table(c.TableName()).Where("router_link = ? ", routerLink).First(&category).Error
	if err != nil {
		return nil, err
	}

	return &category, nil
}
