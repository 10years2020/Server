package userController

import (
	"net/http"

	"github.com/swsad-dalaotelephone/Server/models/task"
	"github.com/swsad-dalaotelephone/Server/models/user"
	"github.com/swsad-dalaotelephone/Server/modules/log"
	"github.com/swsad-dalaotelephone/Server/modules/util"

	"github.com/gin-gonic/gin"
)

/*
GetPublishedTasks : get published task
require: cookie
return: publsihed task list
*/
func GetPublishedTasks(c *gin.Context) {

	// publisherId := c.Query("publisher_id")
	user := c.MustGet("user").(userModel.User)
	publisherId := user.Id

	// get published tasks
	tasks, err := taskModel.GetTasksByStrKey("publisher_id", publisherId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "can not fetch task list",
		})
		log.ErrorLog.Println(err)
		c.Error(err)
		return
	}

	tasksJson, err := util.StructToJsonStr(tasks)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "json convert error",
		})
		log.ErrorLog.Println(err)
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"tasks": tasksJson,
	})
	log.InfoLog.Println(publisherId, len(tasks), "success")
}
