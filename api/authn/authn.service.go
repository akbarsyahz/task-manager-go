package authn

import (
	"taskManager/db"
	dtoM "taskManager/db/model"
)

func CreateUser() {
	database := db.Connection()
	user := &dtoM.User{NameFirst: "Jinzhu", NameLast: "Suntzu",Age: 18, DateOfBirth: "2020-01-20", PlaceBirth: "Mana Aja"}
	result := database.Create(user)
	err := result.Error
	if err != nil {
		panic(err)
	}
}