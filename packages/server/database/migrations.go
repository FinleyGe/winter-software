package database

import (
	"SelectionSystem/model"

	"gorm.io/gorm"
)

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.User{},
		&model.Advice{},
		&model.Form{},
		&model.Avatar{},
		&model.Reason{},
	)
}
