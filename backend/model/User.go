package model

import (
	"cloud-note/config"
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
	result := db.NewRecord(user)
	if result {
		db.Create(user)
	}
    return result
}

func (user *User) Find(username string) {
	db := config.GetDB()
	db.Where("username = ?", username).First(user)
}