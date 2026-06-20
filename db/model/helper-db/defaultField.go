package helperdb

import (
	"time"

	"github.com/google/uuid"
)

type UserBy struct {
	CreatedBy uuid.UUID `json:"created_by"`
	UpdatedBy uuid.UUID `json:"updated_by"`
	DeletedBy uuid.UUID `json:"deleted_by"`
}

type DefaultField struct {
	ID        uuid.UUID `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index"`
}

type DefaultFieldUser struct {
	DefaultField
	UserBy
}
