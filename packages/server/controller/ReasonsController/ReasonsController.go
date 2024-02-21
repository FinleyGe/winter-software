package ReasonsController

import (
	"SelectionSystem/database"
	"SelectionSystem/model"
	"SelectionSystem/service/ReasonsService"
	"SelectionSystem/utility"
	"github.com/gin-gonic/gin"
)

type ShowReasonsData struct {
	Account string `json:"account"`
}

type DeleteReasonData struct {
	Value int `json:"value"`
}

// 显示理由
func ShowReasons(c *gin.Context) {
	var data ShowReasonsData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utility.JsonResponseInternalServerError(c)
		return
	}
	err = database.DB.Where("account = ?", data.Account).First(&model.Reason{}).Error
	if err != nil {
		utility.JsonResponse(404, "当前没有理由", nil, c)
		return
	}
	rc := make([]model.Reason, 0)
	err = database.DB.Where("account = ?", data.Account).Find(&rc).Error
	if err != nil {
		utility.JsonResponseInternalServerError(c)
		return
	}
	utility.JsonSuccessResponse(c, rc)
}

// 添加理由
func AddReasons(c *gin.Context) {
	var s model.Reason
	err := c.ShouldBindJSON(&s)
	if err != nil {
		utility.JsonResponseInternalServerError(c)
		return
	}
	err = ReasonsService.Add(&s)
	if err != nil {
		utility.JsonResponseInternalServerError(c)
		return
	}
	utility.JsonSuccessResponse(c, nil)
}

// 删除理由
func DeleteReasons(c *gin.Context) {
	var data DeleteReasonData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utility.JsonResponseInternalServerError(c)
		return
	}
	err = ReasonsService.Delete(data.Value)
	if err != nil {
		utility.JsonResponseInternalServerError(c)
		return
	}
	utility.JsonSuccessResponse(c, nil)
}
