package StudentController

import (
	"SelectionSystem/service/StudentService"
	"SelectionSystem/utility"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func UpdateAvatar(c *gin.Context) {
	var account string
	account = c.PostForm("account")

	form, err := c.MultipartForm()
	if err != nil {
		utility.JsonResponseInternalServerError(c)
		return
	}

	applyForm, err := StudentService.QueryApplyFormByAccount(account)

	//删除文件
	err = os.Remove(applyForm.AvatarAddress)
	if err != nil {
		if err != nil {
			fmt.Println("Error:", err)
			utility.JsonResponse(500, "文件删除失败", nil, c)
			return
		}
	} else {
		// 如果文件成功删除，打印成功信息
		fmt.Println("File successfully deleted")
	}

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
