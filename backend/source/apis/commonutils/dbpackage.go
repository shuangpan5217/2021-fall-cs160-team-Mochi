package commonutils

import "time"

type User struct {
	Username  string     `gorm:"primary_key;not null" json:"username"`
	Password  string     `gorm:"not null" json:"password"`
	CreatedAt *time.Time `gorm:"default:CURRENT_TIMESTAMP;" json:"created_at"`
	UpdatedAt *time.Time `gorm:"default:CURRENT_TIMESTAMP;" json:"updated_at"`
}
