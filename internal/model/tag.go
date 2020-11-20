package model

import (
	"time"

	"github.com/kekaiwang/go-blog/pkg/drives"
)

type Tag struct {
	Id         int64     `gorm:"primary_key"`
	Name       string    `gorm:"name" json:"name"`               //名称
	RouterLink string    `gorm:"router_link" json:"router_link"` // 路由链接
	UseTimes   int64     `gorm:"use_times" json:"use_times"`     // tag链接文章数量
	Status     int       `gorm:"status" json:"status"`           // 1:正常 2:禁用
	CreatedAt  time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt  time.Time `gorm:"updated_at" json:"updated_at"`
}

var tag *Tag

func TagModel() *Tag {
	return tag
}

func (t *Tag) TableName() string {
	return `tag`
}

// Create insert tag
func (t *Tag) Create(tag *Tag) error {
	return drives.BlogDB.Create(tag).Error
}

func (t *Tag) GetTagByIds(ids string) ([]*Tag, error) {
	tags := []*Tag{}
	err := drives.BlogDB.Table(t.TableName()).Where("id in (?) ", ids).Find(&tags).Error
	if err != nil {
		return nil, err
	}

	return tags, nil
}
