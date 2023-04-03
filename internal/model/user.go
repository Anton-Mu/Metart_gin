package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	ID          uint   `json:"id" gorm:"column:user_id"`
	Username    string `json:"username" gorm:"column:user_name"`
	Email       string `json:"email" gorm:"column:email"`
	Password    string `json:"-" gorm:"column:password"`
	PaymentInfo string `json:"payment_info" gorm:"column:payment_info"`
}

// TableName 数据库表名自定义，默认为model的复数形式，比如这里默认为 users
func (User) TableName() string {
	return "users"
}
