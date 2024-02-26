package routes

import (
	"net/http"

	"example.com/myproject/api/middlewares"
	"example.com/myproject/api/models"
	"github.com/gin-gonic/gin"
)

func Updates(c *gin.Context) {
	db, err := middlewares.DbConnection()
	if err != nil {
		panic(err)
	}
	id := c.Param("id")

	var member_update models.Members

	if err := c.BindJSON(&member_update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := models.Members{}

	if err := db.Model(&updates).Where("Id = ?", id).Updates(map[string]interface{}{"Name": member_update.Name}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "update succesfull")

}
