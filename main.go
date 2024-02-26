package main

import (
	"os"

	"example.com/myproject/api/routes"
	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"

)

func init() {
	gotenv.Load()
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	port := os.Getenv("PORT")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/signup", routes.SignIn)
	r.DELETE("/delete/:id", routes.Delete)
	r.PUT("/update/:id", routes.Updates)
	r.GET("all_members", routes.Get_all)
	r.GET("get_one/:id", routes.Get_one_users)
	r.Run(":" +port)
}
