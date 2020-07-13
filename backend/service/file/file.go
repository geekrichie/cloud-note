package file

import (
	"cloud-note/model"
)

func GetFolderFirstFile(folderid int) model.File{
	var file model.File
	model.GetFirstFileByFolderId(folderid,&file)
	return file
}
