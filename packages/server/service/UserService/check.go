package UserService

import (
	"SelectionSystem/database"
	"SelectionSystem/model"
)

// 判断账号是否存在，存在返回密码和身份
func CheckAccountExistAndBackPwdAndIdentity(account string) (int, string, string, error) {
	var user model.User
	err := database.DB.Where("account=?", account).First(&user).Error
	if err != nil {
		return 0, "", "", err
	} else {
		return user.NumOfStudent, user.Password, user.Identity, nil
	}
}

// 判断密码是否正确
func CheckPwd(truePassword string, password string) bool {
	if truePassword == password {
		return true
	} else {
		return false
	}
}
