package folder

import (
	"cloud-note/model"
	"cloud-note/service/file"
	"cloud-note/service/folder"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetFoldersAction(c *gin.Context) {
     uid, existed := c.Get("uid")
     if existed == false {
     	log.Panic("uid is not existed")
	 }
	 var folders []model.Folder
	if value, ok := uid.(int); ok {
		folder.GetUserFolders(value, &folders)
	}
	var res []interface{}
	for _,fd := range folders {
		item := make(map[string]interface{})
        item["folder_id"] = fd.ID
        item["name"] = fd.Name
        re := file.GetFolderLastFile(fd.ID)
		item["files"] = "";
		if !re.EqualObject() {
			item["files"] = re.ID
		}
        res = append(res,item)
	}
	c.JSON(http.StatusOK,gin.H{
		"code" : 1,
		"data" : res,
	})
}

func GetFilesAction(c *gin.Context) {
	folderid := c.Query("folderid")
	if folderid == "" {
		c.JSON(http.StatusOK,gin.H{
			"code" : 0,
			"data" : nil,
		})
		return
	}
	fd,_ := strconv.Atoi(folderid)
	var files []model.File
    file.GetFilesByFolderId(fd,&files)
	c.JSON(http.StatusOK,gin.H{
		"code" : 1,
		"data" : files,
	})
}
