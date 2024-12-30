package controllers

import (
	"go-project/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

type CreatePostInput struct {
	UserName string `json:"username" binding:"required"`
	
}
type UpdatePostInput struct {
	UserName string `json:"username"`
}

func CreateUser(c *gin.Context){
	var input CreatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post := models.User{Username: input.UserName}
	models.DB.Create(&post)
	c.JSON(http.StatusOK, gin.H{"data": post})
}
func FindUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func FindUser(c *gin.Context){
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "User not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func UpdateUser(c *gin.Context){
	var input UpdatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "User not found!"})
		return
	}
	models.DB.Model(&user).Updates(models.User{Username: input.UserName})
	c.JSON(http.StatusOK, gin.H{"data": user})
}
func DeleteUser(c *gin.Context){
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "User not found!"})
		return
	}
	models.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"data": "success"})
}