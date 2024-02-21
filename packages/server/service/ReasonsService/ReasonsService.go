package ReasonsService

import (
	"SelectionSystem/database"
	"SelectionSystem/model"
)

func Add(s *model.Reason) (err error) {
	result := database.DB.Omit("value").Create(&s)
	return result.Error
}

func Delete(value int) (err error) {
	result := database.DB.Delete(&model.Reason{}, value)
	return result.Error
}
