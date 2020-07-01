package model

import (
	"cloud-note/config"
)

type User struct {
	Username string  `gorm:"column:username" `
	Password string
	Email   string  `gorm:"column:email"`
	InsertTime string
	UpdateTime string
}

func (user *User) Save() bool{
	db := config.GetDB()
	db.Create(user)
	result := db.NewRecord(user)
    return result
}

func (user *User) Find(username string) {
	db := config.GetDB()
	db.Where("username = ?", username).First(user)
}