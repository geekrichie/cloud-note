package folder

import (
	"cloud-note/model"
)

func GetUserFolders(uid int, folders *[]model.Folder){
	var fd model.Folder
	fd.GetFoldersByUid(uid, folders)
}