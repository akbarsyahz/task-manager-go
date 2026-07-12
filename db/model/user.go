package model

import (
	helperdb "taskManager/db/model/helper-db"
	"time"

	"github.com/google/uuid"
)

// User represents a user in the system.
type User struct {
	helperdb.DefaultFieldUser
	NameFirst   string    `json:"name_first"`
	NameLast    string    `json:"name_last"`
	Age         uint      `json:"age"`
	DateOfBirth time.Time `json:"date_of_birth"`
	PlaceBirth  string    `json:"place_birth"`
	Tasks       []Task
	UserRoles   []UserRole `gorm:"foreignKey:UserID"`
}

// UserRole defines the field(s) used to look up a Junction table for two table user + role.
type UserRole struct {
	UserID uuid.UUID `gorm:"primaryKey"`
	RoleID uuid.UUID `gorm:"primaryKey"`

	User User `gorm:"foreignKey:UserID"`
	Role Role `gorm:"foreignKey:RoleID"`
}

// Role defines the field(s) used to look up a Junction table for UserRole.
type Role struct {
	ID   uuid.UUID
	Name string
}

// LoginUser defines the field(s) used to look up a Junction table for authentication user.
type LoginUser struct {
	UserID   uuid.UUID `gorm:"primaryKey"`
	User     User      `gorm:"foreignKey:UserID;references:ID"`
	Username string    `json:"username"`
	Hash     []byte    `json:"hash"`
	Salted   []byte    `json:"salted"`
}
