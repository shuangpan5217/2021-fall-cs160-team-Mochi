package commonutils

import "github.com/jinzhu/gorm"

func InitDatabase(db *gorm.DB) error {
	err := db.AutoMigrate(&User{}).Error
	return err
}
