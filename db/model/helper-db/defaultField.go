package helperdb

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UserBy defines the field(s) used to look up a user.
type UserBy struct {
	CreatedBy uuid.UUID `json:"created_by"`
	UpdatedBy uuid.UUID `json:"updated_by"`
	DeletedBy uuid.UUID `json:"deleted_by"`
}

// DefaultField defines the field(s) used to look up a custom default.
type DefaultField struct {
	ID        uuid.UUID `gorm:"type:uuid;primarykey;default:gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// DefaultFieldUser defines the field(s) used to look up a combined two custome struct.
type DefaultFieldUser struct {
	DefaultField
	UserBy
}
