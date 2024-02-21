package TFormService

import (
	"SelectionSystem/database"
	"SelectionSystem/model"
	"SelectionSystem/service/TeacherService"
	"time"
)

func TUpdateForm(form *model.Form, state int, advice string, account string) error {
	change := 0
	if state == 1 && (form.State == 0 || form.State == 2) {
		change = 1
	}
	if state == 2 && form.State == 1 {
		change = -1
	}
	err := TeacherService.UpdateNum(account, change)
	if err != nil {
		return err
	}
	err = database.DB.Model(form).Update("state", state).Error
	if err != nil {
		return err
	}
	err = database.DB.Model(form).Update("time_of_teacher", time.Now().Format("2006-01-02 15:04:05")).Error
	if err != nil {
		return err
	}
	return database.DB.Model(form).Update("advice", advice).Error
}
