package StudentService

import (
	"SelectionSystem/database"
	"SelectionSystem/model"
)

// 通过账号查询申报表单
func QueryApplyFormByAccount(account string) (model.Form, error) {
	//判断是否提交过表单
	err := database.DB.Where("account=?", account).First(&model.Form{}).Error
	if err != nil {
		return model.Form{}, err
	}

	var applyForm model.Form
	database.DB.Where("account=?", account).Find(&applyForm)
	return applyForm, nil
}
