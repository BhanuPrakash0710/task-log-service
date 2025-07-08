package controllers

import (
	"net/http"

	"github.com/BhanuPrakash0710/to-do-list-api/internal/utils"
	"github.com/BhanuPrakash0710/to-do-list-api/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Register(c *gin.Context) {
	var input models.RegisterInputUser
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		ID:       primitive.NewObjectID().Hex(),
		Email:    input.Email,
		Name:     input.Name,
		Password: utils.HashPassword(input.Password),
		//UserId:   utils.GenerateUserId(input.Name),
	}

	if err := models.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func Login(c *gin.Context) {
	var input models.LoginInputUser
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := models.GetUserByEmail(input.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if !utils.VerifyPassword(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	// Generate JWT token
	jwtToken, err := utils.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.SetCookie("jwt_token", jwtToken, 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"Id":      user.ID,
		"token":   jwtToken,
	})
}
