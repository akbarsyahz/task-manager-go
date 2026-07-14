package task

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateTaskHandler godoc
// @Summary Create Task
// @Tags task
// @Accept json
// @Produce json
// @Param body body CreateTaskDto true "Task"
// @Success 200 {object} map[string]interface{}
// @Router /task/ [post]
// CreateTaskHandler Handle Task Create
func CreateTaskHandler(ctx *gin.Context) {
	var formTask CreateTaskDto

	if err := ctx.ShouldBind(&formTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := CreateTask(formTask)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Success",
	})
}

// UpdateTaskHandler godoc
// @Summary Create Task
// @Tags task
// @Accept json
// @Produce json
// @Param body body UpdateTaskDto true "Task"
// @Success 200 {object} map[string]interface{}
// @Router /task/ [patch]
// UpdateTaskHandler Handle Task Update
func UpdateTaskHandler(ctx *gin.Context) {
	var formTask UpdateTaskDto
	if err := ctx.ShouldBind(&formTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, errUser := ctx.Get("sub")
	if !errUser {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	_, err := UpdateTask(formTask, userID.(uuid.UUID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Success",
	})
}

// GetTaskHandler godoc
// @Summary Create Task
// @Tags task
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /task/all [get]
// GetTaskHandler Handle Task Update
func GetTaskHandler(ctx *gin.Context) {
	userID, errUser := ctx.Get("sub")
	if !errUser {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	_, errGetTask := GetAllTask(userID.(uuid.UUID))
	if errGetTask != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": errGetTask.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Success",
	})
}

// GetTaskDetailHandler godoc
// @Summary Create Task
// @Tags task
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Success 200 {object} map[string]interface{}
// @Router /task/{id} [get]
// GetTaskDetailHandler Handle Task Update
func GetTaskDetailHandler(ctx *gin.Context) {
	taskID := ctx.Query("id")
	parsedUUID, err := uuid.Parse(taskID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result, err := GetTaskDetail(parsedUUID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Success",
		"date":    result,
	})
}

// DeleteTaskDetailHandler godoc
// @Summary Delete Task
// @Tags task
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Success 200 {object} map[string]interface{}
// @Router /task/{id} [delete]
func DeleteTaskDetailHandler(ctx *gin.Context) {
	// var taskId []uuid.UUID
	taskID := ctx.Query("id")
	parsedUUID, err := uuid.Parse(taskID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result, err := DeleteTask(parsedUUID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Success",
		"date":    result,
	})
}
