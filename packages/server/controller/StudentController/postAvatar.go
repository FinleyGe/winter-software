package StudentController

import (
	"SelectionSystem/model"
	"SelectionSystem/service/StudentService"
	"SelectionSystem/utility"
	"github.com/gin-gonic/gin"
	"os"
)

func PostAvatar(c *gin.Context) {
	var account string
	account = c.PostForm("account")

	form, err := c.MultipartForm()
	if err != nil {
		utility.JsonResponseInternalServerError(c)
		return
	}

	applyForm, err := StudentService.QueryApplyFormByAccount(account)
	var empty model.Form
	if applyForm == empty {
		applyForm.Account = account
		_, err := StudentService.CreateApplyForm(&applyForm)
		if err != nil {
			utility.JsonResponseInternalServerError(c)
			return
		}
	}

	//创建文件夹
	path := "./storeFile/Avatar/" + account
	utility.Mkdir(path)
	files := form.File["file"]

	//存储文件并将路径保存到数据库
	for _, file := range files {
		err := c.SaveUploadedFile(file, path+"/"+file.Filename)
		if err != nil {
			utility.JsonResponse(400, "文件上传失败", nil, c)
			return
		}

		dir, _ := os.Getwd()
		applyForm.AvatarAddress = dir + "/storeFile/Avatar/" + account + "/" + file.Filename
		err = StudentService.UpdateForm(applyForm)
		if err != nil {
			utility.JsonResponseInternalServerError(c)
			return
		}
	}
	utility.JsonResponse(200, "OK", nil, c)
}
