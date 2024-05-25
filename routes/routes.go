package routes

import (
	"fmt"
	"os"
	"sync"

	"example.com/myproject/config"
	"example.com/myproject/handlers"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
}

func wrapHandler(handler func(c *gin.Context, wg *sync.WaitGroup), wg *sync.WaitGroup) gin.HandlerFunc {
	return func(c *gin.Context) {
		wg.Add(1)
		handler(c, wg)
	}
}

func setupRoutes(r *gin.Engine, wg *sync.WaitGroup) {
	r.POST("/signup", wrapHandler(handlers.SignUp, wg))
	r.GET("/users", wrapHandler(handlers.GetUsers, wg))
	r.GET("/user/:id", wrapHandler(handlers.GetUser, wg))
	r.PUT("/user/:id", wrapHandler(handlers.UpdateUser, wg))
	r.DELETE("/user/:id", wrapHandler(handlers.DeleteUser, wg))
}

func Run() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	wg := &sync.WaitGroup{}

	setupRoutes(r, wg)

	port := os.Getenv("PORT")
	err := r.Run(":" + port)

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("\t\t\tServer running..")
	}
}
