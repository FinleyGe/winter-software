package TeacherService

import (
	"SelectionSystem/database"
	"SelectionSystem/model"
)

func QueryNum(account string) int {
	var user model.User
	database.DB.Model(&user).Where("account =?", account)
	return user.NumOfStudent
}

func UpdateNum(account string, change int) error {
	var user model.User
	database.DB.Model(&user).Where("account =?", account)
	num := user.NumOfStudent
	num += change
	err := database.DB.Model(&user).Where("account =?", account).Update("num_of_student", num).Error
	return err
}
