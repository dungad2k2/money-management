package main

import (
	"go-project/models"
    "go-project/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()  
    router.LoadHTMLGlob("views/*")
    router.GET("/", func (c *gin.Context) {
        c.HTML(200, "index.html", gin.H{
            "title": "Home Page",
        })
    })
 
    models.ConnectDB() 
    router.POST("/users", controllers.CreateUser)
    router.GET("/users", controllers.FindUsers)
    router.GET("/users/:id", controllers.FindUser)
    router.PATCH("/users/:id", controllers.UpdateUser)
    router.DELETE("/users/:id", controllers.DeleteUser)
    router.LoadHTMLGlob("views/*")
    router.GET("/users/:id/expenses", func(c *gin.Context) {
        c.HTML(200, "expense.html", gin.H{
            "title": "Expenses Page",
            "user_id": c.Param("user_id"),
        })
    })

    router.POST("/users/:id/expenses", controllers.CreateExpense)
    router.GET("/users/:id/expenses/all", controllers.FindExpenses)
    router.GET("/expenses/:id", controllers.FindExpense)
    router.PATCH("/expenses/:id", controllers.UpdateExpense)
    router.DELETE("/expenses/:id", controllers.DeleteExpense)
    router.Run(":8080")
}