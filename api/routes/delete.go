package routes

import (
	"net/http"

	"example.com/myproject/api/middlewares"
	"example.com/myproject/api/models"
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	db, err := middlewares.DbConnection()
	if err != nil {
		panic(err)
	}

	member_delete := &models.Members{}
	id := c.Param("id")

	if err := db.Delete(member_delete, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	

}
