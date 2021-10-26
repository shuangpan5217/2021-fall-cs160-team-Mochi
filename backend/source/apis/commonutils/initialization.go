package commonutils

import (
	"2021-fall-cs160-team-Mochi/backend/source/apis/dbpackages"

	"github.com/jinzhu/gorm"
)

func InsertTables(db *gorm.DB) error {
	err := db.AutoMigrate(&dbpackages.User{}).Error
	if err == nil {
		err = db.AutoMigrate(&dbpackages.Group{}).Error
	}
	if err == nil {
		err = db.AutoMigrate(dbpackages.GroupUser{}).Error
	}
	if err == nil {
		err = db.AutoMigrate(dbpackages.Note{}).Error
	}
	if err == nil {
		err = db.AutoMigrate(dbpackages.UserNote{}).Error
	}
	if err == nil {
		err = db.AutoMigrate(dbpackages.GroupNote{}).Error
	}
	if err == nil {
		err = db.AutoMigrate(dbpackages.Comment{}).Error
	}
	if err == nil {
		err = db.AutoMigrate(dbpackages.Friend{}).Error
	}
	return err
}

func AddFKConstraints(db *gorm.DB) {
	dbpackages.CreateFKConstraints(db)
}
