package dbpackages

import "time"

type User struct {
	Username    string     `gorm:"primary_key;not null" json:"username"`
	Password    string     `gorm:"not null" json:"password"`
	Email       string     `json:"email,omitempty"`
	FirstName   string     `json:"first_name,omitempty"`
	MiddleName  string     `json:"middle_name,omitempty"`
	LastName    string     `json:"last_name,omitempty"`
	Description string     `json:"description,omitempty"`
	CreatedAt   *time.Time `gorm:"default:CURRENT_TIMESTAMP;" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"default:CURRENT_TIMESTAMP;" json:"updated_at"`
}
