package folder

import (
	"cloud-note/model"
)

func GetUserFolders(uid int) ([]model.Folder){
	folders := model.GetFoldersByUid(uid)
	return folders
}