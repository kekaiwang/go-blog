package drives

import (
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/kekaiwang/go-blog/config"
)

var (
	BlogDB *gorm.DB
)

func BlogDBInit() {
	initBlogDB()    // init DB
	initBlogImgDB() // init img DB
}

// layze init
// init db
func initBlogDB() {
	var (
		conf       = config.Get()
		blogDBOnce sync.Once
		initDB     func()
	)

	// init DB
	initDB = func() {
		db, err := gorm.Open("mysql", conf.Mysql.Uri)
		if err != nil {
			panic(err)
		}

		// blog db
		BlogDB = db // db
		BlogDB.DB().SetMaxOpenConns(conf.Mysql.MaxOpenConns)
		BlogDB.DB().SetMaxIdleConns(conf.Mysql.MaxIdleConns)
		BlogDB.DB().SetConnMaxLifetime(time.Minute * time.Duration(conf.Mysql.ConnMaxLifeTime))
		BlogDB.LogMode(conf.Mysql.LogMode)
	}

	if BlogDB == nil {
		blogDBOnce.Do(initDB) // init once
	} else if err := BlogDB.DB().Ping(); err != nil {
		initDB()
	} else if err := BlogDB.Error; err != nil {
		panic(err)
	}
}

// layze init
// initBlogImgDB db
func initBlogImgDB() {
	var (
		conf      = config.Get()
		imgDBOnce sync.Once
		iniImgDB  func()
	)

	// init DB
	iniImgDB = func() {
		db, err := gorm.Open("mysql", conf.Mysql.Uri)
		if err != nil {
			panic(err)
		}

		BlogDB = db
		BlogDB.DB().SetMaxOpenConns(conf.Mysql.MaxOpenConns) // max connect
		BlogDB.DB().SetMaxIdleConns(conf.Mysql.MaxIdleConns)
		BlogDB.DB().SetConnMaxLifetime(time.Minute * time.Duration(conf.Mysql.ConnMaxLifeTime))
		BlogDB.LogMode(conf.Mysql.LogMode)
	}

	if BlogDB == nil {
		imgDBOnce.Do(iniImgDB)
	} else if err := BlogDB.DB().Ping(); err != nil {
		iniImgDB()
	} else if err := BlogDB.Error; err != nil {
		panic(err)
	}
}
