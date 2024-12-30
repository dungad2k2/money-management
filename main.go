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
    router.PUT("/users/:id", controllers.UpdateUser)
    router.DELETE("/users/:id", controllers.DeleteUser)
    router.POST("/expenses", controllers.CreateExpense)
    router.GET("/expenses", controllers.FindExpenses)
    router.GET("/expenses/:id", controllers.FindExpense)
    router.PUT("/expenses/:id", controllers.UpdateExpense)
    router.DELETE("/expenses/:id", controllers.DeleteExpense)
    
    router.Run(":8080")
}