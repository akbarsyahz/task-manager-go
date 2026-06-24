package authn

import (
	"taskManager/db"
	dtoM "taskManager/db/model"

	"gorm.io/gorm"
)

func CreateUser(inputUser dtoM.User, inputLogin dtoM.LoginUser) (string , error) {
	database := db.Connection()
	user := &dtoM.User{NameFirst: inputUser.NameFirst, NameLast: inputUser.NameLast, Age: inputUser.Age, DateOfBirth: inputUser.DateOfBirth, PlaceBirth: inputUser.PlaceBirth}
	loginUser := &dtoM.LoginUser{Username: inputLogin.Username, Password: inputLogin.Password}
	result := database.Transaction(func(tx *gorm.DB) error {
		if err := database.Create(user).Error; err != nil {
			return err
		}
		
		if err := database.Create(loginUser).Error; err != nil {
			return err
		}
		
		return nil
	})
	
	return "Success", result
}