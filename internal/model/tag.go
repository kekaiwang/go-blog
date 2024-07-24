package model

import (
	"time"

	"github.com/kekaiwang/go-blog/pkg/drives"
)

type Tag struct {
	ID         int64     `gorm:"id" json:"id"`
	Name       string    `gorm:"name" json:"name"`               //名称
	RouterLink string    `gorm:"router_link" json:"router_link"` // 路由链接
	UseTimes   int64     `gorm:"use_times" json:"use_times"`     // tag链接文章数量
	Status     int       `gorm:"status" json:"status"`           // 1:正常 2:禁用
	Created    time.Time `gorm:"created" json:"created"`         // created time
	Updated    time.Time `gorm:"updated" json:"updated"`         // update time
}

var tag *Tag

func TagModel() *Tag {
	return tag
}

func (t *Tag) TableName() string {
	return `tag`
}

// Create insert tag
func (t *Tag) Create() error {
	return drives.BlogDB.Create(t).Error
}

// GetTagByIds.
func (t *Tag) GetTagByIds(ids []int) ([]*Tag, error) {
	tags := []*Tag{}
	err := drives.BlogDB.Table(t.TableName()).Where("id in (?) ", ids).Find(&tags).Error
	if err != nil {
		return nil, err
	}

	return tags, nil
}

// GetTagByQuery.
func (t *Tag) GetTagByQuery(query string, args []interface{}) (*Tag, error) {
	tag := &Tag{}
	err := drives.BlogDB.Table(t.TableName()).Where(query, args...).First(&tag).Error
	if err != nil {
		return nil, err
	}

	return tag, nil
}

// GetTagByRouterLink.
func (t *Tag) GetTagByRouterLink(routerLink string) (*Tag, error) {
	tag := Tag{}

	err := drives.BlogDB.Table(t.TableName()).Where("router_link = ? ", routerLink).First(&tag).Error
	if err != nil {
		return nil, err
	}

	return &tag, nil
}

// GetTagList.
func (t *Tag) GetTagList(query string, args []interface{}, limit, offset int64) ([]*Tag, error) {
	var tags []*Tag
	err := drives.BlogDB.Table(t.TableName()).Where(query, args...).Limit(limit).Offset(offset).Order("id DESC").Find(&tags).Error
	if err != nil {
		return nil, err
	}

	return tags, nil
}

// CountTag.
func (t *Tag) CountTag(query string, args []interface{}) (int64, error) {
	var total int64 // total num
	err := drives.BlogDB.Table(t.TableName()).Where(query, args...).Count(&total).Error
	if err != nil {
		return total, err
	}

	return total, nil
}

// UpdateTag.
func (t *Tag) UpdateTag() (int64, error) {
	result := drives.BlogDB.Model(&t).Update(t)
	if result.Error != nil {
		return 0, result.Error
	}

	return int64(result.RowsAffected), nil
}
