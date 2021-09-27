package commonutils

import (
	"2021-fall-cs160-team-Mochi/backend/source/apis/dbpackages"

	"github.com/jinzhu/gorm"
)

func InitDatabase(db *gorm.DB) error {
	err := db.AutoMigrate(&dbpackages.User{}).Error
	return err
}
