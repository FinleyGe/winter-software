package UserService

import (
	"SelectionSystem/database"
	"SelectionSystem/model"
)

func UpdatePwd(account string, newPassword string) {
	database.DB.Model(&model.User{}).Where("account=?", account).Update("password", newPassword)
}
