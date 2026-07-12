package task

import "github.com/google/uuid"

// CreateTaskDto defines the field(s) used to look up a Table.
type CreateTaskDto struct {
	UserID      uuid.UUID `form:"user_id" binding:"required"`
	Title       string    `form:"title" binding:"required"`
	Status      string    `form:"status" binding:"required"`
	Description string    `form:"description" binding:"required"`
}

// UpdateTaskDto defines the field(s) used to look up a Table.
type UpdateTaskDto struct {
	ID uuid.UUID `form:"id" binding:"required"`
	// UserID      uuid.UUID `form:"user_id" binding:"required"`
	Title       string `form:"title" binding:"required"`
	Status      string `form:"status" binding:"required"`
	Description string `form:"description" binding:"required"`
}
