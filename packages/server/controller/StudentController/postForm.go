package StudentController

import (
	"SelectionSystem/model"
	"SelectionSystem/service/StudentService"
	"SelectionSystem/utility"
	"github.com/gin-gonic/gin"
	"strconv"
)

func PostApplyForm(c *gin.Context) {
	var data model.Form
	data.Name = c.PostForm("name")
	data.Account = c.PostForm("account")
	score := c.PostForm("score")
	data.Score, _ = strconv.ParseFloat(score, 64)
	data.PhoneNumber = c.PostForm("phone_number")
	data.Position = c.PostForm("position")
	rank := c.PostForm("rank")
	data.Rank, _ = strconv.Atoi(rank)
	math1 := c.PostForm("math_of_1")
	data.MathOf1, _ = strconv.Atoi(math1)
	math2 := c.PostForm("math_of_2")
	data.MathOf2, _ = strconv.Atoi(math2)
	cpp1 := c.PostForm("cpp_of_1")
	data.CppOf1, _ = strconv.Atoi(cpp1)
	cpp2 := c.PostForm("cpp_of_2")
	data.CppOf2, _ = strconv.Atoi(cpp2)
	data.Profile = c.PostForm("profile")
	data.TeacherName = c.PostForm("teacher_name")

	//将申报表单存入数据库
	_, err := StudentService.CreateApplyForm(&data)
	if err != nil {
		utility.JsonResponseInternalServerError(c)
		return
	}
	//
	//form, err := c.MultipartForm()
	//if err != nil {
	//	utility.JsonResponseInternalServerError(c)
	//	return
	//}
	//
	////创建文件夹
	//path := "./storeFile/" + id
	//utility.Mkdir(path)
	//files := form.File["file"]
	//var fileInfo model.FileInfo
	//fileInfo.FormId = id
	//
	////存储文件并将路径保存到数据库
	//for _, file := range files {
	//	err := c.SaveUploadedFile(file, path+"/"+file.Filename)
	//	if err != nil {
	//		utility.JsonResponse(400, "文件上传失败", nil, c)
	//		return
	//	}
	//
	//	fileInfo.Address = "127.0.0.1:8080/storeFile/" + id + "/" + file.Filename
	//	err = StudentService.CreateFileInfo(&fileInfo)
	//	if err != nil {
	//		utility.JsonResponseInternalServerError(c)
	//		return
	//	}
	//}
	utility.JsonResponse(200, "OK", nil, c)
}
