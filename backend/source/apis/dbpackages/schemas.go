package dbpackages

import (
	"time"

	"github.com/google/uuid"
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
)

// create foreign key constraints
const (
	// users and groups
	group_users_username_fk = `ALTER TABLE group_users DROP CONSTRAINT IF EXISTS group_user_username_fk; ALTER TABLE group_users ADD CONSTRAINT group_user_username_fk FOREIGN KEY (username) REFERENCES Users(username) ON DELETE CASCADE;`
	group_users_groupID_fk  = `ALTER TABLE group_users DROP CONSTRAINT IF EXISTS group_users_groupID_fk; ALTER TABLE group_users ADD CONSTRAINT group_users_groupID_fk FOREIGN KEY (group_id) REFERENCES Groups(group_id) ON DELETE CASCADE;`

	// users and notes
	user_notes_username_fk = `ALTER TABLE user_notes DROP CONSTRAINT IF EXISTS user_notes_username_fk; ALTER TABLE user_notes ADD CONSTRAINT user_notes_username_fk FOREIGN KEY (username) REFERENCES Users(username) ON DELETE CASCADE;`
	user_notes_noteID_fk   = `ALTER TABLE user_notes DROP CONSTRAINT IF EXISTS user_notes_noteID_fk; ALTER TABLE user_notes ADD CONSTRAINT user_notes_noteID_fk FOREIGN KEY (note_id) REFERENCES Notes(note_id) ON DELETE CASCADE;`

	// groups and notes
	group_notes_groupID_fk = `ALTER TABLE group_notes DROP CONSTRAINT IF EXISTS group_notes_groupID_fk; ALTER TABLE group_notes ADD CONSTRAINT group_notes_groupID_fk FOREIGN KEY (group_id) REFERENCES Groups(group_id) ON DELETE CASCADE;`
	group_notes_noteID_fk  = `ALTER TABLE group_notes DROP CONSTRAINT IF EXISTS group_notes_noteID_fk; ALTER TABLE group_notes ADD CONSTRAINT group_notes_noteID_fk FOREIGN KEY (note_id) REFERENCES Notes(note_id) ON DELETE CASCADE;`

	// notes and comments
	note_comments_noteID_fk = `ALTER TABLE comments DROP CONSTRAINT IF EXISTS note_comments_noteID_fk; ALTER TABLE comments ADD CONSTRAINT note_comments_noteID_fk FOREIGN KEY (note_id) REFERENCES Notes(note_id) ON DELETE CASCADE`
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
	GroupID     uuid.UUID `gorm:"primary_key;not null" json:"group_id"`
	GroupName   string    `gorm:"not null" json:"group_name"` // unique for the same user
	Description string    `json:"description"`
	GroupOwner  string    `gorm:"not null" json:"group_owner"` // username of the user
}

type GroupUser struct {
	GroupID  uuid.UUID `gorm:"not null" json:"group_id"` // once the group is deleted, all related group users will be deleted
	Username string    `gorm:"not null" json:"username"` // once the user is deleted, all related group users will be deleted
}

type Note struct {
	NoteID        uuid.UUID `gorm:"primary_key;not null" json:"note_id"`
	NoteOwner     string    `gorm:"not null" json:"note_owner"` // username of note owner. Once delted, all notes will be deleted
	Description   string    `json:"description"`
	Title         string    `json:"title"`
	Category      string    `json:"category"`
	NoteReference string    `gorm:"not null" json:"note_reference"`
	Type          string    `json:"type"`
	Tag           string    `gorm:"not null" json:"tag"`
}

// users have access to notes
type UserNote struct {
	Username string    `gorm:"not null" json:"username"`
	NoteID   uuid.UUID `gorm:"not null" json:"note_id"`
}

// groups have access to notes
type GroupNote struct {
	GroupID uuid.UUID `gorm:"not null" json:"group_id"`
	NoteID  uuid.UUID `gorm:"not null" json:"note_id"`
}

type Comment struct {
	CommentID uuid.UUID `gorm:"primary_key;not null" json:"comment_id"`
	NoteID    uuid.UUID `gorm:"not null" json:"note_id"` // once deleted, all comments will be deleted
	Username  string    `gorm:"not null" json:"username"`
	Content   string    `gorm:"not null" json:"content"`
}

func CreateFKConstraints(db *gorm.DB) {
	// users and groups
	db.Exec(group_users_username_fk)
	db.Exec(group_users_groupID_fk)

	// users and notes
	db.Exec(user_notes_username_fk)
	db.Exec(user_notes_noteID_fk)

	// groups and notes
	db.Exec(group_notes_groupID_fk)
	db.Exec(group_notes_noteID_fk)

	// notes and comments
	db.Exec(note_comments_noteID_fk)
}
