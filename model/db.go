package model

import (
	"fmt"
	"myblog/utils"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func InitDb() {
	db, err = gorm.Open(mysql.Open(
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			utils.Conf.User,
			utils.Conf.Password,
			utils.Conf.Host,
			utils.Conf.Port,
			utils.Conf.DbName)),
		&gorm.Config{})
	if err != nil {
		fmt.Println("gorm.Open failed, err : ", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("db.DB failed, err : ", err)
	}

	err = db.AutoMigrate(&User{}, &Article{}, &Category{})
	if err != nil {
		fmt.Println("db.AutoMigrate failed, err : ", err)
	}

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	//defer sqlDB.Close()

}
