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

func (f *File)GetFirstFileByFolderId(folderid int) {
	db := config.GetDB()
	db.Where("folder_id = ?", folderid).First(f)
}
func (f *File)GetLastFileByFolderId(folderid int) {
	db := config.GetDB()
	db.Where("folder_id = ?", folderid).Last(f)
}
func (f *File)GetAllByUid(files *[]File ) {
	db := config.GetDB()
	db.Where("user_id = ?", f.UserId).Find(files)
}
func (f *File)GetAllByFolderId(files *[]File ) {
	db := config.GetDB()
	db.Where("folder_id = ?", f.FolderId).Find(files)
}


func (f File)EqualObject() bool{
	if f == (File{}) {  // 括号不能去
		return true;
	}
	return false
}

func (f *File)GetFilesByFolderId(files *[]File ) {
	db := config.GetDB()
	db.Where("folder_id = ?", f.UserId).Find(files)
}