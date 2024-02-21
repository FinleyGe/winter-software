package StudentService

import (
	"SelectionSystem/database"
	"SelectionSystem/model"
)

func UpdateForm(data model.Form) (err error) {
	var form model.Form
	form.Name = data.Name
	form.Account = data.Account
	form.Score = data.Score
	form.PhoneNumber = data.PhoneNumber
	form.Position = data.Position
	form.Rank = data.Rank
	form.MathOf1 = data.MathOf1
	form.MathOf2 = data.MathOf2
	form.CppOf1 = data.CppOf1
	form.CppOf2 = data.CppOf2
	form.Profile = data.Profile
	form.AvatarAddress = data.AvatarAddress
	form.TeacherName = data.TeacherName
	form.AutographAddressOfTeacher = data.AutographAddressOfTeacher
	form.AutographAddressOfStudent = data.AutographAddressOfStudent
	var beforeForm model.Form
	err = database.DB.Where("account = ?", data.Account).First(&beforeForm).Error
	if err != nil {
		return err
	}
	err = database.DB.Model(&beforeForm).Omit("id").Updates(form).Error
	return err
}
