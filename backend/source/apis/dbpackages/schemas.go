package dbpackages

import (
	"time"

	"github.com/jinzhu/gorm"
)

const (
	// table names
	UserTable      = "users"
	GroupTable     = "groups"
	GroupUserTable = "group_users"
	NoteTable      = "notes"
	UserNoteTable  = "user_notes"
	GroupNoteTable = "group_notes"
	CommentTable   = "comments"
)

const (
	// create foreign key constraints
	group_users_username_fk = `ALTER TABLE group_users DROP CONSTRAINT IF EXISTS group_user_username_fk; ALTER TABLE group_users ADD CONSTRAINT group_user_username_fk FOREIGN KEY (username) REFERENCES Users(username) ON DELETE CASCADE;`
	group_users_groupID_fk  = `ALTER TABLE group_users DROP CONSTRAINT IF EXISTS group_users_groupID_fk; ALTER TABLE group_users ADD CONSTRAINT group_users_groupID_fk FOREIGN KEY (group_id) REFERENCES Groups(group_id) ON DELETE CASCADE;`
)

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

type Group struct {
	GroupID string `gorm:"primary_key;not null" json:"group_id"`
}

type GroupUser struct {
	Username string `gorm:"not null" json:"username"`
	GroupID  string `gorm:"not null" json:"group_id"`
}

type Note struct {
}

type UserNote struct {
}

type GroupNote struct {
}

type Comment struct {
}

func CreateFKConstraints(db *gorm.DB) {
	db.Exec(group_users_username_fk)
	db.Exec(group_users_groupID_fk)
}
