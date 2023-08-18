package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"wedding-invitation-api/user/entities"
)

var DB *gorm.DB
var err error

func init() {
	dsn := "root:ming199951@tcp(47.111.131.243:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	// AutoMigrate 用于自动迁移您的 schema，保持您的 schema 是最新的。如数据库添加字段，它会自动维护最新的数据库表
	// 这句代码会自动创建表结构，但是数据库必须手动创建
	err = DB.AutoMigrate(&entities.User{})
	if err != nil {
		return
	}
}
