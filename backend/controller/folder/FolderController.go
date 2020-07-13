package folder

import (
	"cloud-note/model"
	"cloud-note/service/file"
	"cloud-note/service/folder"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetFoldersAction(c *gin.Context) {
     uid, existed := c.Get("uid")
     if existed == false {
     	log.Panic("uid is not existed")
	 }
	 var folders []model.Folder
	if value, ok := uid.(int); ok {
		folders = folder.GetUserFolders(value)
		fmt.Println(folders)
	}
	var res []interface{}
	for _,fd := range folders {
		item := make(map[string]interface{})
        item["folder_id"] = fd.ID
        item["name"] = fd.Name
        re := file.GetFolderFirstFile(fd.ID)
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
