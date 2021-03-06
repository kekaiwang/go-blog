package model

import (
	"time"

	"github.com/kekaiwang/go-blog/pkg/drives"
)

//Article 文章
type AdminUser struct {
	ID         int       `ggorm:"id" json:"id"`
	Name       string    `gorm:"name" json:"name"`
	Password   string    `gorm:"password" json:"password"`
	Email      string    `gorm:"email" json:"email"`
	LoginCount int       `gorm:"login_count" json:"login_count"`
	Salt       string    `gorm:"salt" json:"salt"`
	Status     int8      `gorm:"status" json:"status"`
	LastLogin  time.Time `gorm:"last_login" json:"last_login"`
	LastIp     string    `gorm:"last_ip" json:"last_ip"`
	Created    time.Time `gorm:"created" json:"created"`
	Updated    time.Time `gorm:"updated" json:"updated"`
}

var user *AdminUser

func UserModel() *AdminUser {
	return user
}

func (u *AdminUser) TableName() string {
	return `admin_user`
}

// GetUser. get admin user
func (u *AdminUser) GetUser(query string, args []interface{}) (*AdminUser, error) {
	var user AdminUser
	err := drives.BlogDB.Table(u.TableName()).Where(query, args...).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *AdminUser) UpdateLogin() (int, error) {
	result := drives.BlogDB.Model(&u).Update(u)
	if err := result.Error; err != nil {
		return 0, err
	}

	return int(result.RowsAffected), nil
}
