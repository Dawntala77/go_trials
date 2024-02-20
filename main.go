package main

import (
	"example.com/myproject/api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/signup", routes.SignIn)
	r.DELETE("/delete/:id", routes.Delete)
	r.PUT("/update/:id", routes.Updates)
	r.Run(":8000")
}
