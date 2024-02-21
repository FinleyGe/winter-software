package UserController

import (
	"SelectionSystem/service/UserService"
	"SelectionSystem/utility"
	"github.com/gin-gonic/gin"
	"log"
)

type LoginData struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func LoginCheck(c *gin.Context) {
	var data LoginData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Println(err)
		utility.JsonResponseInternalServerError(c)
		return
	}

	//检查账号是否存在并返回密码和身份
	num, password, identity, err := UserService.CheckAccountExistAndBackPwdAndIdentity(data.Account)
	if err != nil {
		utility.JsonResponse(404, "账号不存在", nil, c)
		return
	}

	//判断密码是否正确
	flag := UserService.CheckPwd(password, data.Password)
	if !flag {
		utility.JsonResponse(405, "密码错误", nil, c)
		return
	}

	//返回账号信息
	utility.JsonSuccessResponse(c, gin.H{
		"account":        data.Account,
		"identity":       identity,
		"num_of_student": num,
	})
}
