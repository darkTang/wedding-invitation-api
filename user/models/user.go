package models

type User struct { // 默认表名为users
	Id     string `json:"id" gorm:"type:bigint;primaryKey;autoIncrement"`
	Name   string `json:"name" form:"name" gorm:"type:varchar(255);not null"`
	Number uint   `json:"number" form:"number" gorm:"type:tinyint;unsigned;not null"`
	Mobile string `json:"mobile" form:"mobile" gorm:"type:varchar(255);not null"`
}

func (User) TableName() string { // 可以修改默认的表名
	return "user"
}
