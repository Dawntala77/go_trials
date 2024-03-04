package routes

import (
	"example.com/myproject/handlers"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func init() {
	
}

func setupRoutes(r *gin.Engine) {
	r.POST("/signup", handlers.SignUp)
	r.GET("/users", handlers.GetUsers)
	r.GET("/user/:id", handlers.GetUser)
	r.PUT("/user/:id", handlers.UpdateUser)
	r.DELETE("/user/:id", handlers.DeleteUser)
}

func Run() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	
	setupRoutes(r)

	port := os.Getenv("PORT")
	err := r.Run(":" + port)
	
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("\t\t\tServer running..")
	}
}
