package model

import "cloud-note/config"

type Folder struct{
	ID int
	UserID   int
	Name string
	UpdateTime string
    CreateTime string
}

func (f Folder) TableName() string {
	return "folder"
}

func (fd *Folder)GetFoldersByUid(uid int,folders *[]Folder){
	db := config.GetDB()
	db.Where("user_id = ?", uid).Find(&folders)
}