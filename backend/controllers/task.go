package controllers

import (
	"net/http"
	"taskido/backend/database"
	"taskido/backend/models"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {

	var tasks []models.Task
	database.DB.Find(&tasks)

	c.JSON(http.StatusOK, gin.H{"data": tasks})

}

func CreateTask(c *gin.Context) {
	var input models.Task

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := models.Task{
		TaskId:      input.TaskId,
		Title:       input.Description,
		Description: input.Description,
		Priority:    input.Priority,
		Status:      input.Status,
		DueDate:     input.DueDate,
		Tag:         input.Tag,
		CreatedAt:   input.CreatedAt,
		UpdatedAt:   input.UpdatedAt,
	}

	database.DB.Create(&task)

	c.JSON(http.StatusOK, gin.H{"Data: ": task})
}

func UpdateTask(c *gin.Context) {
	var input models.Task
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var task models.Task

	if err := database.DB.Where("task_id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task not Found"})
		return
	}

	database.DB.Model(&task).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": task})
}

func GetTask(c *gin.Context) {
	var task models.Task

	if err := database.DB.Where("task_id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

func DeleteTask(c *gin.Context) {

	var task models.Task

	if err := database.DB.Where("task_id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task Not Found"})
		return
	}

	database.DB.Delete(task)
	c.JSON(http.StatusOK, gin.H{"data": task})

}
