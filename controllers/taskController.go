package controllers

import (
	"strconv"
	"time"

	"github.com/BhanuPrakash0710/to-do-list-api/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTask(c *gin.Context) {
	var input models.InputTask
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Get userID from context
	userID, exists := c.Get("ID")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	now := time.Now()

	task := models.Task{
		ID:          primitive.NewObjectID().Hex(),
		UserID:      userID.(string),
		Title:       input.Title,
		Status:      input.Status,
		Description: input.Description,
		CreatedAt:   now.Format("2006-01-02 15:04:05"),
		UpdatedAt:   now.Format("2006-01-02 15:04:05"),
	}

	if err := models.AddOneTask(&task); err != nil {
		c.JSON(500, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(201, task)
}

func GetAllTasks(c *gin.Context) {
	// Get userID from context
	userID, exists := c.Get("ID")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	pageQuery := c.DefaultQuery("page", "1")
	perPageQuery := c.DefaultQuery("perPage", "5")

	page, _ := strconv.ParseInt(pageQuery, 10, 64)
	perPage, _ := strconv.ParseInt(perPageQuery, 10, 64)

	// Fallback values
	if page <= 0 {
		page = 1
	}
	if perPage <= 0 {
		perPage = 5
	}
	tasks, err := models.GetTasks(userID.(string), page, perPage)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve tasks"})
		return
	}
	c.JSON(200, tasks)
}

func GetTaskByID(c *gin.Context) {
	taskID := c.Param("id")

	// Get userID from context
	_, exists := c.Get("ID")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	task, err := models.GetOneTask(taskID)
	if err != nil {
		c.JSON(404, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(200, task)
}

func UpdateTaskByID(c *gin.Context) {
	taskID := c.Param("id")

	var input models.PatchTask
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// Get userID from context
	_, exists := c.Get("ID")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	update := bson.M{}
	if input.Title != nil {
		update["title"] = *input.Title
	}
	if input.Description != nil {
		update["description"] = *input.Description
	}
	if input.Status != nil {
		update["status"] = *input.Status
	}

	if len(update) == 0 {
		c.JSON(400, gin.H{"error": "No fields to update"})
		return
	}

	update["updatedAt"] = time.Now().Format("2006-01-02 15:04:05")
	err := models.UpdateOneTask(taskID, update)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update task"})
		return
	}
	c.JSON(200, gin.H{"message": "Task updated successfully"})
}

func DeleteTaskByID(c *gin.Context) {
	taskID := c.Param("id")

	// Get userID from context
	_, exists := c.Get("ID")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	err := models.DeleteOneTask(taskID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete task"})
		return
	}
	c.JSON(200, gin.H{"message": "Task deleted successfully"})
}
