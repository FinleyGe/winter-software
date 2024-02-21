package StudentController

import (
	"SelectionSystem/service/StudentService"
	"SelectionSystem/utility"
	"github.com/gin-gonic/gin"
	"log"
)

type QueryFormData struct {
	Account string `json:"account"`
}

// 查询表单
func QueryApplyForm(c *gin.Context) {
	var data QueryFormData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Println(err)
		utility.JsonResponseInternalServerError(c)
		return
	}

	//根据账号查询表单并返回
	applyForm, err := StudentService.QueryApplyFormByAccount(data.Account)
	if err != nil {
		utility.JsonResponse(404, "未提交过申报", nil, c)
		return
	}
	utility.JsonSuccessResponse(c, applyForm)
}
