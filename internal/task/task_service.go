package task

import (
	"fmt"

	model "github.com/akbarsyahz/task-manager-go/db/model"

	"github.com/google/uuid"
)

// CreateTaskService making Task
func CreateTaskService(input CreateTaskDto) (string, error) {
	tasks := model.Task{
		UserID:      input.UserID,
		Title:       input.Title,
		Status:      input.Status,
		Description: input.Description,
	}

	_, err := CreateTask(tasks)

	if err != nil {
		return "Failed", fmt.Errorf("%v", err.Error())
	}

	return "Success", nil
}

// UpdateTaskService for updateing the task by input user
func UpdateTaskService(input UpdateTaskDto, userID uuid.UUID) (string, error) {
	tasks := model.Task{
		// NOTE: (akbar): userId Already input by set
		Title:       input.Title,
		Status:      input.Status,
		Description: input.Description,
	}
	tasks.ID = input.ID
	tasks.UserID = userID

	if _, err := UpdateTask(tasks); err != nil {
		return "Failed", fmt.Errorf("failed to update task: %w", err)
	}

	return "Success", nil
}

// GetAllTaskService for getting all task, but its must be specified by user task
func GetAllTaskService(userID uuid.UUID) ([]model.Task, error) {
	result, err := GetTask(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get tasks: %w", err)
	}

	return result, nil
}

// GetDetailTaskService getting task detail
func GetDetailTaskService(taskID uuid.UUID) (model.Task, error) {
	result, err := GetDetailTask(taskID)
	if err != nil {
		return result, err
	}
	return result, nil
}

// DeleteTaskService deleting task
func DeleteTaskService(taskID uuid.UUID) (string, error) {
	_, err := DeleteTask(taskID)
	if err != nil {
		return "error:", err
	}
	return "Success", nil
}
