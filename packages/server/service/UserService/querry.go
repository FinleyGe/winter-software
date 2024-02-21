package UserService

import (
	"SelectionSystem/database"
	"SelectionSystem/model"
)

// 查询账号返回密码
func GetPwdByAccount(account string) string {
	var user model.User
	database.DB.Where("account=?", account).First(&user)
	return user.Password
}
