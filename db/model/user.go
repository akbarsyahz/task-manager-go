package model

import (
	helperdb "taskManager/db/model/helper-db"
)

type User struct {
	helperdb.DefaultFieldUser
	NameFirst   string `json:"name_first"`
	NameLast    string `json:"name_last"`
	Age         uint   `json:"age"`
	DateOfBirth string `json:"date_of_birth"`
	PlaceBirth  string `json:"place_birth"`
	Tasks       []Task
}
