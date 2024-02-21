package SignatureController

import (
	"SelectionSystem/database"
	"SelectionSystem/model"
	"SelectionSystem/service/StudentService"
	"SelectionSystem/utility"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
)

func CreateSignature(c *gin.Context) {
	var identity string
	identity = c.PostForm("identity")
	var id int
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		utility.JsonResponseInternalServerError(c)
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		utility.JsonResponseInternalServerError(c)
		return
	}

	var applyForm model.Form
	err = database.DB.Where("id = ?", id).First(&applyForm).Error
	if err != nil {
		utility.JsonResponse(404, "表单不存在", nil, c)
		return
	}

	path := "./storeFile/Signature/" + strconv.Itoa(id) + "/"
	if identity == "student" {
		path += "Student"
	}
	if identity == "teacher" {
		path += "Teacher"
	}
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
		dir += "/storeFile/Signature/" + strconv.Itoa(id) + "/"
		if identity == "student" {
			applyForm.AutographAddressOfStudent = dir + "Student/" + file.Filename
		}
		if identity == "teacher" {
			applyForm.AutographAddressOfTeacher = dir + "Teacher/" + file.Filename
		}
		err = StudentService.UpdateForm(applyForm)
		if err != nil {
			utility.JsonResponseInternalServerError(c)
			return
		}
	}
	utility.JsonResponse(200, "OK", nil, c)
}
