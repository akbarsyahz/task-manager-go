package model

import (
	helperdb "taskManager/db/model/helper-db"
)

type Task struct {
	helperdb.DefaultFieldUser
	UserId      uint   `json:"user_id"`
	Title       string `json:"title"`
	Status      bool   `json:"status"`
	Description string `json:"description"`
}
