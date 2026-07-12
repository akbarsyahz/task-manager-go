package repositoryauthn

import (
	"fmt"
	"taskManager/db"
	dtoM "taskManager/db/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// GetUser for getting username and return the user result
func GetUser(username string) (dtoM.LoginUser, error) {
	var user dtoM.LoginUser
	database := db.Connection()

	result := database.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("username = ?", username).First(&user).Error; err != nil {
			return err
		}
		return nil
	})

	if result != nil {
		return user, result
	}

	return user, nil
}

// GetUserProfile Getting profile user: Name, birth and etc
func GetUserProfile(id uuid.UUID) (dtoM.User, error) {
	var user dtoM.User
	database := db.Connection()

	result := database.Transaction(func(tx *gorm.DB) error {
		if err := tx.Find(&user, id).Error; err != nil {
			return err
		}
		return nil
	})

	if result != nil {
		return user, fmt.Errorf("%s", result.Error())
	}

	return user, nil
}

// CreateUser Getting profile user: Name, birth and etc
func CreateUser(user *dtoM.User, identityLogin *dtoM.LoginUser) error {
	database := db.Connection()
	result := database.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return err
		}

		if err := tx.Create(identityLogin).Error; err != nil {
			return err
		}
		return nil
	})

	return result
}
