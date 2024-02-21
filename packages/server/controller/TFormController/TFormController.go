package TFormController

import (
	"SelectionSystem/database"
	"SelectionSystem/model"
	"SelectionSystem/service/TFormService"
	"SelectionSystem/utility"
	"github.com/gin-gonic/gin"
)

type ShowFormData struct {
	TeacherName string `json:"teacher_name"`
}

type HandelFormData struct {
	Id      int    `json:"id"`
	State   int    `json:"state"`
	Advice  string `json:"advice"`
	Account string `json:"account"`
}

func ShowForm(c *gin.Context) {
	var data ShowFormData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utility.JsonResponseInternalServerError(c)
		return
	}
	var form model.Form
	err = database.DB.Where("teacher_name = ?", data.TeacherName).First(&form).Error
	if err != nil {
		utility.JsonResponse(404, "当前没有表单", nil, c)
		return
	}
	var ac = make([]model.Form, 0)
	err = database.DB.Where("teacher_name = ?", data.TeacherName).Find(&ac).Error
	utility.JsonSuccessResponse(c, ac)
}

func HandelForm(c *gin.Context) {
	var data HandelFormData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utility.JsonResponseInternalServerError(c)
		return
	}
	var form model.Form
	err = database.DB.Where("id = ?", data.Id).First(&form).Error
	if err != nil {
		utility.JsonResponse(404, "表单不存在", nil, c)
		return
	}
	err = TFormService.TUpdateForm(&form, data.State, data.Advice, data.Account)
	if err != nil {
		utility.JsonResponseInternalServerError(c)
		return
	}
	utility.JsonSuccessResponse(c, nil)
}
