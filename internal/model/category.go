package model

import (
	"time"

	"github.com/kekaiwang/go-blog/pkg/drives"
)

// Category category model
type Category struct {
	ID          int64     `gorm:"id" json:"id"`
	Name        string    `gorm:"name" json:"name"`                 //名称
	RouterLink  string    `gorm:"router_link" json:"router_link"`   // 路由链接
	LinkArticle int64     `gorm:"link_article" json:"link_article"` // 链接文章数量
	Status      int       `gorm:"status" json:"status"`             // 1:正常 2:禁用 3:已删除
	Created     time.Time `gorm:"created_at" json:"created"`        // created time
	Updated     time.Time `gorm:"updated_at" json:"updated"`        // updated time
}

var category *Category

func CategoryModel() *Category {
	return category
}

func (c *Category) TableName() string {
	return `category`
}

// Create insert category
func (c *Category) Create() error {
	return drives.BlogDB.Create(c).Error
}

// GetAll. get all category
func (c *Category) GetAll() ([]*Category, error) {
	categories := []*Category{}

	err := drives.BlogDB.Table(c.TableName()).Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}

// GetAll get all by where
func (c *Category) GetAllByWhere(query string, args []interface{}) ([]*Category, error) {
	categories := []*Category{}

	err := drives.BlogDB.Table(c.TableName()).Where(query, args...).Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}

// GetCategoryByRouterLink get category by router link info
func (c *Category) GetCategoryByRouterLink(routerLink string) (*Category, error) {
	category := Category{}

	err := drives.BlogDB.Table(c.TableName()).Where("router_link = ? ", routerLink).First(&category).Error
	if err != nil {
		return nil, err
	}

	return &category, nil
}

// GetCategoryList.
func (c *Category) GetCategoryList(query string, args []interface{}, limit, offset int64) ([]*Category, error) {
	var categories []*Category

	err := drives.BlogDB.Table(c.TableName()).Where(query, args...).Order("id DESC").Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}

// CountCategory. count category
func (c *Category) CountCategory(query string, args []interface{}) (int64, error) {
	var total int64

	err := drives.BlogDB.Table(c.TableName()).Where(query, args).Count(&total).Error
	if err != nil {
		return 0, err
	}

	return total, nil
}

// GetCategory. get catetory
func (c *Category) GetCategory(query string, args []interface{}) (*Category, error) {
	category := &Category{}
	err := drives.BlogDB.Table(c.TableName()).Where(query, args...).First(&category).Error
	if err != nil {
		return nil, err
	}

	return category, nil
}

// UpdateCategory. update category
func (c *Category) UpdateCategory() (int64, error) {
	result := drives.BlogDB.Model(&c).Update(c)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}

// UpdateCategory. update category
func (c *Category) UpdateCategoryBy() (int64, error) {
	result := drives.BlogDB.Model(&c).Update(c)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}

// UpdateCategory. update category
func (c *Category) UpdateCategorys() (int64, error) {
	result := drives.BlogDB.Model(&c).Update(c)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}

// UpdateCategory. update category
func (c *Category) UpdateCategoryInfo() (int64, error) {
	result := drives.BlogDB.Model(&c).Update(c)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}
