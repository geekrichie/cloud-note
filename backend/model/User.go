package model

import (
	"cloud-note/config"
	"fmt"
)

type User struct {
	Username string  `gorm:"column:username;unique;not null" `
	Password string  `gorm:"column:password;not null"`
	Email   string  `gorm:"column:email;unique;not null"`
	InsertTime string
	UpdateTime string
}

func (user *User) Save() bool{
	db := config.GetDB()
	if err := db.Create(user).Error; err != nil {
		fmt.Println(err);
		return false;
	}
	return true;
}

func (user *User) Find(username string) {
	db := config.GetDB()
	db.Where("username = ?", username).First(user)
}