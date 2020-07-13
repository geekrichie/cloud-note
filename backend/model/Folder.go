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

func GetFoldersByUid(uid int) ([]Folder){
	db := config.GetDB()
	var folders []Folder
	db.Where("user_id = ?", uid).Find(&folders)
	return folders
}