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
	// 这句代码会自动创建表结构，但是数据库必须手动创建
	if err := DB.AutoMigrate(&entities.User{}); err != nil {
		return
	}
}
