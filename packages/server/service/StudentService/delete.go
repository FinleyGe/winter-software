package StudentService

import (
	"SelectionSystem/database"
	"SelectionSystem/model"
)

func DeleteForm(account string) error {
	err := database.DB.Where("account=?", account).Delete(model.Form{}).Error
	return err
}
