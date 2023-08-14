package entities

type User struct { // 默认表名为users
	Id     string `json:"id" gorm:"type:int unsigned;primaryKey;autoIncrement"`
	Name   string `json:"name" gorm:"type:varchar(255);not null"`
	Number uint   `json:"number" gorm:"type:tinyint unsigned;unsigned;not null"`
	Mobile string `json:"mobile" gorm:"type:varchar(255);not null"`
}

func (User) TableName() string { // 可以修改默认的表名
	return "user"
}
