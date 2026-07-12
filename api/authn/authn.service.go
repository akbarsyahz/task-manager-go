package authn

import (
	"fmt"
	repo "taskManager/api/authn/repository"
	"taskManager/api/helper-api/security"
	"taskManager/db"
	dtoM "taskManager/db/model"
	"time"
)

// Registration for registration user
func Registration(inputUser UserRegisterDto) (string, error) {
	layoutDate := "2006-01-02"
	convertToValidDate, errDate := time.Parse(layoutDate, inputUser.DateOfBirth)
	if errDate != nil {
		return "failed", fmt.Errorf("something error when generate")
	}

	resultSalt, err := security.GenerateSalt()
	if err != nil {
		return "failed", fmt.Errorf("something error when generate")
	}

	password, errHash := security.HashArgon2Register(inputUser.Password, resultSalt)
	if errHash != nil {
		return "", errHash
	}

	//nolint:lll
	user := &dtoM.User{NameFirst: inputUser.NameFirst, NameLast: inputUser.NameLast, Age: inputUser.Age, DateOfBirth: convertToValidDate, PlaceBirth: inputUser.PlaceBirth}
	loginUser := &dtoM.LoginUser{Username: inputUser.Username, Hash: password, Salted: resultSalt}

	errorCreating := repo.CreateUser(user, loginUser)
	if errorCreating != nil {
		return "Failed", fmt.Errorf("failed register user")
	}

	return "Success", errorCreating
}

// Login User input username and password and return the token
func Login(usernameInput string, passwordInput string) (string, error) {
	user, err := repo.GetUser(usernameInput)
	if err != nil {
		return "", fmt.Errorf("failed to get user: %w", err)
	}

	_, errComparing := security.ComparingPassword(passwordInput, user)
	if errComparing != nil {
		return "", fmt.Errorf("username or password is wrong: %w", errComparing)
	}

	profile, err := repo.GetUserProfile(user.UserID)
	if err != nil {
		return "", fmt.Errorf("failed to get profile: %w", err)
	}

	token, err := security.CreatingToken(profile)
	if err != nil {
		return "", fmt.Errorf("failed to creating token: %w", err)
	}

	return token, nil
}

// GetAllUser return all user
func GetAllUser() ([]dtoM.User, error) {
	database := db.Connection()
	var users []dtoM.User
	result := database.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
