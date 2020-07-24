package file

import (
	"cloud-note/model"
)

func GetFolderFirstFile(folderid int) model.File{
	var file model.File
	file.GetFirstFileByFolderId(folderid)
	return file
}
func GetFolderLastFile(folderid int) model.File{
	var file model.File
	file.GetLastFileByFolderId(folderid)
	return file
}

func GetFilesByUid(uid int,files *[]model.File) {
    file := model.File{}
    file.UserId = uid
    file.GetAllByUid(files)
}
func GetFilesByFolderId(folderid int,files *[]model.File) {
	file := model.File{}
	file.FolderId = folderid
	file.GetAllByFolderId(files)
}