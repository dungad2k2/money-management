package controllers
import (
	"go-project/models"
	"net/http"
	"github.com/gin-gonic/gin"
)
type CreateExpenseInput struct {
	UserName string `json:"username" binding:"required"`
	Amount float64 `json:"amount" binding:"required"`
	Reason string `json:"reason" binding:"required"`
}
type UpdateExpenseInput struct {
	UserName string `json:"username"`
	Amount float64 `json:"amount"`
	Reason string `json:"reason"`
}

func CreateExpense(c *gin.Context){
	var input CreateExpenseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	if err := models.DB.Where("username = ?", input.UserName).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "User not found!"})
		return
	}
	expense := models.Expense{UserName: input.UserName, Amount: input.Amount, Reason: input.Reason}
	models.DB.Create(&expense)
	user.Moneyspend += input.Amount
	models.DB.Save(&user)
	c.JSON(http.StatusOK, gin.H{"data": expense})
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
	var input UpdateExpenseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var expense models.Expense
	if err := models.DB.Where("id = ?", c.Param("id")).First(&expense).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Expense not found!"})
		return
	}
	var user models.User
	if err := models.DB.Where("username = ?", input.UserName).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "User not found!"})
		return
	}
	user.Moneyspend -= expense.Amount
	if input.Amount != 0 {
		expense.Amount = input.Amount
		user.Moneyspend += input.Amount
	}
	if input.Reason != "" {
		expense.Reason = input.Reason
	}
	models.DB.Save(&expense)
	models.DB.Save(&user)
	c.JSON(http.StatusOK, gin.H{"data": expense})
}
func DeleteExpense(c *gin.Context) {
    var expense models.Expense
    if err := models.DB.Where("id = ?", c.Param("id")).First(&expense).Error; err != nil {
        c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Expense not found!"})
        return
    }

    var user models.User
    if err := models.DB.Where("username = ?", expense.UserName).First(&user).Error; err != nil {
        c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "User not found!"})
        return
    }

    user.Moneyspend -= expense.Amount
    models.DB.Save(&user)
    models.DB.Delete(&expense)

    c.JSON(http.StatusOK, gin.H{"data": "success"})
}