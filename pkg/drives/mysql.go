package drives

import (
	"sync"
	"time"

	//import mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/kekaiwang/go_blog/config"
)

var (
	BlogDB *gorm.DB
)

func BlogDBInit() {
	initBLogDB()
}

func initBLogDB() {
	var (
		conf       = config.Get()
		blogDBOnce sync.Once
	)

	initDB := func() {
		db, err := gorm.Open("mysql", conf.Mysql.Uri)
		if err != nil {
			panic(err)
		}

		BlogDB = db
		BlogDB.DB().SetMaxOpenConns(conf.Mysql.MaxOpenConns)
		BlogDB.DB().SetMaxIdleConns(conf.Mysql.MaxIdleConns)
		BlogDB.DB().SetConnMaxLifetime(time.Minute * time.Duration(conf.Mysql.ConnMaxLifeTime))
		BlogDB.LogMode(conf.Mysql.LogMode)
	}

	if BlogDB == nil {
		blogDBOnce.Do(initDB)
	} else if err := BlogDB.DB().Ping(); err != nil {
		initDB()
	} else if err := BlogDB.Error; err != nil {
		panic(err)
	}
}
