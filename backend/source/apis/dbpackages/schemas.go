package dbpackages

import (
	"time"

	"github.com/jinzhu/gorm"
)

const (
	// table names iteral
	UserTable      = "users"
	GroupTable     = "groups"
	GroupUserTable = "group_users"
	NoteTable      = "notes"
	UserNoteTable  = "user_notes"
	GroupNoteTable = "group_notes"
	CommentTable   = "comments"
	FriendTable    = "friends"
)

// create foreign key constraints
const (
	// users and groups access
	group_users_username_fk = `ALTER TABLE group_users DROP CONSTRAINT IF EXISTS group_user_username_fk; ALTER TABLE group_users ADD CONSTRAINT group_user_username_fk FOREIGN KEY (username) REFERENCES Users(username) ON DELETE CASCADE;`
	group_users_groupID_fk  = `ALTER TABLE group_users DROP CONSTRAINT IF EXISTS group_users_groupID_fk; ALTER TABLE group_users ADD CONSTRAINT group_users_groupID_fk FOREIGN KEY (group_id) REFERENCES Groups(group_id) ON DELETE CASCADE;`

	// users and groups owner fk
	group_users_groupOwner_fk = `ALTER TABLE groups DROP CONSTRAINT IF EXISTS group_users_groupOwner_fk; ALTER TABLE groups ADD CONSTRAINT group_users_groupOwner_fk FOREIGN KEY (group_owner) REFERENCES Users(username) ON DELETE CASCADE;`

	// users and notes owner fk
	user_notes_noteOwner_fk = `ALTER TABLE notes DROP CONSTRAINT IF EXISTS user_notes_noteOwner_fk; ALTER TABLE notes ADD CONSTRAINT user_notes_noteOwner_fk FOREIGN KEY (note_owner) REFERENCES Users(username) ON DELETE CASCADE;`

	// users and notes access
	user_notes_username_fk = `ALTER TABLE user_notes DROP CONSTRAINT IF EXISTS user_notes_username_fk; ALTER TABLE user_notes ADD CONSTRAINT user_notes_username_fk FOREIGN KEY (username) REFERENCES Users(username) ON DELETE CASCADE;`
	user_notes_noteID_fk   = `ALTER TABLE user_notes DROP CONSTRAINT IF EXISTS user_notes_noteID_fk; ALTER TABLE user_notes ADD CONSTRAINT user_notes_noteID_fk FOREIGN KEY (note_id) REFERENCES Notes(note_id) ON DELETE CASCADE;`

	// groups and notes access
	group_notes_groupID_fk = `ALTER TABLE group_notes DROP CONSTRAINT IF EXISTS group_notes_groupID_fk; ALTER TABLE group_notes ADD CONSTRAINT group_notes_groupID_fk FOREIGN KEY (group_id) REFERENCES Groups(group_id) ON DELETE CASCADE;`
	group_notes_noteID_fk  = `ALTER TABLE group_notes DROP CONSTRAINT IF EXISTS group_notes_noteID_fk; ALTER TABLE group_notes ADD CONSTRAINT group_notes_noteID_fk FOREIGN KEY (note_id) REFERENCES Notes(note_id) ON DELETE CASCADE;`

	// notes and comments
	note_comments_noteID_fk = `ALTER TABLE comments DROP CONSTRAINT IF EXISTS note_comments_noteID_fk; ALTER TABLE comments ADD CONSTRAINT note_comments_noteID_fk FOREIGN KEY (note_id) REFERENCES Notes(note_id) ON DELETE CASCADE`

	// users and friends
	user_friends_username_fk  = `ALTER TABLE friends DROP CONSTRAINT IF EXISTS user_friends_username_fk; ALTER TABLE friends ADD CONSTRAINT user_friends_username_fk FOREIGN KEY (username) REFERENCES Users(username) ON DELETE CASCADE;`
	user_friends_username2_fk = `ALTER TABLE friends DROP CONSTRAINT IF EXISTS user_friends_username2_fk; ALTER TABLE friends ADD CONSTRAINT user_friends_username2_fk FOREIGN KEY (username2) REFERENCES Users(username) ON DELETE CASCADE;`
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
	GroupID     string     `gorm:"primary_key;not null" json:"group_id"`
	GroupName   string     `gorm:"not null" json:"group_name"` // unique for the same user
	Description string     `json:"description"`
	GroupOwner  string     `gorm:"index;not null" json:"group_owner"` // username of the user.
	CreatedAt   *time.Time `gorm:"default:CURRENT_TIMESTAMP;" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"default:CURRENT_TIMESTAMP;" json:"updated_at"`
}

type GroupUser struct {
	GroupID  string `gorm:"index;not null" json:"group_id"` // once the group is deleted, all related group users will be deleted
	Username string `gorm:"index;not null" json:"username"` // once the user is deleted, all related group users will be deleted
}

type Note struct {
	NoteID        string     `gorm:"primary_key;not null" json:"note_id"`
	NoteOwner     string     `gorm:"index;not null" json:"note_owner"` // username of note owner. Once deleted, all notes will be deleted
	Description   string     `json:"description"`
	Title         string     `json:"title"`
	NoteReference string     `gorm:"not null, index" json:"note_reference"`
	Style         string     `gorm:"not null" json:"style"`
	Type          string     `gorm:"not null" json:"type"`
	Tag           string     `gorm:"not null" json:"tag"`
	CreatedAt     *time.Time `gorm:"default:CURRENT_TIMESTAMP;" json:"created_at"`
	UpdatedAt     *time.Time `gorm:"default:CURRENT_TIMESTAMP;" json:"updated_at"`
}

// users have access to notes
type UserNote struct {
	Username string `gorm:"index;not null" json:"username"`
	NoteID   string `gorm:"index;not null" json:"note_id"`
}

// groups have access to notes
type GroupNote struct {
	GroupID string `gorm:"index;not null" json:"group_id"`
	NoteID  string `gorm:"index;not null" json:"note_id"`
}

type Comment struct {
	CommentID string     `gorm:"primary_key;not null" json:"comment_id"`
	NoteID    string     `gorm:"index;not null" json:"note_id"` // once deleted, all comments will be deleted
	Username  string     `gorm:"not null" json:"username"`
	Content   string     `gorm:"not null" json:"content"`
	CreatedAt *time.Time `gorm:"default:CURRENT_TIMESTAMP;" json:"created_at"`
	UpdatedAt *time.Time `gorm:"default:CURRENT_TIMESTAMP;" json:"updated_at"`
}

type Friend struct {
	Username  string `gorm:"index;not null" json:"username"`
	Username2 string `gorm:"index;not null" json:"username2"`
}

func CreateFKConstraints(db *gorm.DB) {
	// users and groups access
	db.Exec(group_users_username_fk)
	db.Exec(group_users_groupID_fk)

	// users and groups owner fk
	db.Exec(group_users_groupOwner_fk)

	// users and notes owner fk
	db.Exec(user_notes_noteOwner_fk)

	// users and notes access
	db.Exec(user_notes_username_fk)
	db.Exec(user_notes_noteID_fk)

	// groups and notes access
	db.Exec(group_notes_groupID_fk)
	db.Exec(group_notes_noteID_fk)

	// notes and comments
	db.Exec(note_comments_noteID_fk)

	// users and friends
	db.Exec(user_friends_username_fk)
	db.Exec(user_friends_username2_fk)
}
