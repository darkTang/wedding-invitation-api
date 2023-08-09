package models

type User struct { // 默认表名为users
	Id     string `json:"id" gorm:"primaryKey;autoIncrement"`
	Name   string `json:"name" form:"name"`
	Number uint   `json:"number" form:"number"`
	Mobile string `json:"mobile" form:"mobile"`
}

func (User) TableName() string { // 可以修改默认的表名
	return "user"
}
