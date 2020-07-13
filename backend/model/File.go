package model

import (
	"cloud-note/config"
)

type File struct {
	ID  int
	UserId int
	FolderId int
	Title string
	Content string
	UpdateTime string
	CreateTime string
}

func (f File) TableName() string {
	return "article"
}

func GetFirstFileByFolderId(folderid int,file *File ) {
	db := config.GetDB()
	db.Where("folder_id = ?", folderid).First(file)
}
func (f *File)GetAllByUid(file *[]File ) {
	db := config.GetDB()
	db.Where("user_id = ?", f.UserId).Find(file)
}

func (f File)EqualObject() bool{
	if f == (File{}) {  // 括号不能去
		return true;
	}
	return false
}