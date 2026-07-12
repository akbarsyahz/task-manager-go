package repository

import (
	"fmt"
	"taskManager/db"
	dtoM "taskManager/db/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Create This using for createing Task
func Create(input dtoM.Task) (string, error) {
	database := db.Connection()

	result := database.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(input).Error; err != nil {
			return err
		}

		return nil
	})

	if result != nil {
		return "Failed", fmt.Errorf("%v", result.Error())
	}
	return "Success", nil
}

// Update Task
func Update(input dtoM.Task) (string, error) {
	database := db.Connection()
	var tasks dtoM.Task
	task := dtoM.Task{
		Title:       input.Title,
		Status:      input.Status,
		Description: input.Description,
	}
	// Taking id from default model helper easy accsess
	task.ID = input.ID
	task.UserID = input.UserID

	result := database.Transaction(func(tx *gorm.DB) error {
		// NOTE: (akbar): this will auto update from gorm because this is gonna be automatically get the id
		tx.Model(&tasks).Updates(task)
		return nil
	})

	if result != nil {
		return "Failed", fmt.Errorf("%v", result.Error())
	}
	return "Success", nil
}

// Get ALL Task
func Get(userID uuid.UUID) ([]dtoM.Task, error) {
	database := db.Connection()
	var tasks []dtoM.Task
	err := database.Where("user_id = ?", userID).Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// GetDetail Task
func GetDetail(_ uuid.UUID) (dtoM.Task, error) {
	database := db.Connection()
	var tasks dtoM.Task
	result := database.Find(&tasks)
	if result.Error != nil {
		return tasks, result.Error
	}
	return tasks, nil
}

// Delete Task
func Delete(taskID uuid.UUID) (string, error) {
	database := db.Connection()
	var tasks dtoM.Task
	// NOTE: (akbar): We not using condition its will delete task
	result := database.Delete(&tasks, taskID)
	if result.Error != nil {
		return "error:", result.Error
	}
	return "Success", nil
}
