package StudentService

import (
	"SelectionSystem/database"
	"SelectionSystem/model"
	"strconv"
	"time"
)

// 将表格存入数据库
func CreateApplyForm(form *model.Form) (string, error) {
	form.State = 0
	form.TimeOfStudent = time.Now().Format("2006-01-02 15:04:05")
	result := database.DB.Omit("id").Create(form)
	if result.Error != nil {
		return "", result.Error
	}
	return strconv.Itoa(form.Id), nil
}

// 将意见存入数据库
func CreateAdvice(advice *model.Advice) error {
	return database.DB.Omit("id").Create(advice).Error
}
