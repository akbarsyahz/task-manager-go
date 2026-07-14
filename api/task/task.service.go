package task

import (
	"fmt"
	repoTask "taskManager/api/task/repository"
	model "taskManager/db/model"

	"github.com/google/uuid"
)

// CreateTask making Task
func CreateTask(input CreateTaskDto) (string, error) {
	tasks := model.Task{
		UserID:      input.UserID,
		Title:       input.Title,
		Status:      input.Status,
		Description: input.Description,
	}

	_, err := repoTask.Create(tasks)

	if err != nil {
		return "Failed", fmt.Errorf("%v", err.Error())
	}

	return "Success", nil
}

// UpdateTask for updateing the task by input user
func UpdateTask(input UpdateTaskDto, userID uuid.UUID) (string, error) {
	tasks := model.Task{
		// NOTE: (akbar): userId Already input by set
		Title:       input.Title,
		Status:      input.Status,
		Description: input.Description,
	}
	tasks.ID = input.ID
	tasks.UserID = userID

	_, err := repoTask.Update(tasks)

	if err != nil {
		return "Failed", fmt.Errorf("%v", err.Error())
	}

	return "Success", err
}

// GetAllTask for getting all task, but its must be specified by user task
func GetAllTask(userID uuid.UUID) ([]model.Task, error) {
	result, err := repoTask.Get(userID)
	if err != nil {
		return result, fmt.Errorf("%v", err.Error())
	}
	return result, nil
}

// GetTaskDetail getting task detail
func GetTaskDetail(taskID uuid.UUID) (model.Task, error) {
	result, err := repoTask.GetDetail(taskID)
	if err != nil {
		return result, err
	}
	return result, nil
}

// DeleteTask deleting task
func DeleteTask(taskID uuid.UUID) (string, error) {
	_, err := repoTask.Delete(taskID)
	if err != nil {
		return "error:", err
	}
	return "Success", nil
}
