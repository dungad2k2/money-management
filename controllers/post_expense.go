package controllers

import (
	"go-project/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

type CreateMoneyInput struct {
	Ammount float64 `json:"ammount" binding:"required"`
	Reason string `json:"reason" binding:"required"`
	UserName string `json:"username" binding:"required"`
}
type UpdateMoneyInput struct {
	Ammount float64 `json:"ammount"`
	Reason string `json:"reason"`
	UserName string `json:"username"`
}

func CreateExpense(c *gin.Context){
	var input CreateMoneyInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post := models.User{Username: input.UserName}
	models.DB.Create(&post)
	c.JSON(http.StatusOK, gin.H{"data": post})
}
func FindExpenses(c *gin.Context) {
	var expenses []models.Expense
	models.DB.Find(&expenses)
	c.JSON(http.StatusOK, gin.H{"data": expenses})
}
func FindExpense(c *gin.Context){
	var expense models.Expense
	if err := models.DB.Where("id = ?", c.Param("id")).First(&expense).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Expense not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": expense})
}
func UpdateExpense(c *gin.Context){
	var expense models.Expense
	if err := models.DB.Where("id = ?", c.Param("id")).First(&expense).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Expense not found!"})
		return
	}
	var input UpdateMoneyInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedExpense := models.Expense{Amount: input.Ammount, Reason: input.Reason}
	models.DB.Model(&expense).Updates(&updatedExpense)
	
	c.JSON(http.StatusOK, gin.H{"data": expense})
}
func DeleteExpense(c *gin.Context){
	var expense models.Expense
	if err := models.DB.Where("id = ?", c.Param("id")).First(&expense).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Expense not found!"})
		return
	}
	models.DB.Delete(&expense)
	c.JSON(http.StatusOK, gin.H{"data": true})
}