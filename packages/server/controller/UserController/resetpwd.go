package UserController

import (
	"SelectionSystem/service/UserService"
	"SelectionSystem/utility"
	"github.com/gin-gonic/gin"
	"log"
)

type ResetPwdData struct {
	Account     string `json:"account"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func ResetPwd(c *gin.Context) {
	var data ResetPwdData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Println(err)
		utility.JsonResponseInternalServerError(c)
		return
	}

	// 查询账户返回密码
	truePwd := UserService.GetPwdByAccount(data.Account)

	// 判断密码是否正确,正确更新密码
	flag := UserService.CheckPwd(truePwd, data.OldPassword)
	if !flag {
		utility.JsonResponse(405, "原密码错误", nil, c)
		return
	} else {
		UserService.UpdatePwd(data.Account, data.NewPassword)
		utility.JsonSuccessResponse(c, nil)
	}
}
