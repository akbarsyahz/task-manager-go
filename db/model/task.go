package model

import (
	helperdb "taskManager/db/model/helper-db"

	"github.com/google/uuid"
)

// Task defines the field(s) used to look up a Task.
type Task struct {
	helperdb.DefaultFieldUser
	UserID      uuid.UUID `json:"user_id"`
	Title       string    `json:"title"`
	Status      string    `json:"status"`
	Description string    `json:"description"`
}
