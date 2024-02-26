package routes

import (
	"net/http"

	"example.com/myproject/api/middlewares"
	"example.com/myproject/api/models"
	"github.com/gin-gonic/gin"
)

func Get_all(c* gin.Context) {
	db, err := middlewares.DbConnection()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	var all_members []models.Members
	data := db.Find(&all_members)

	if data.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	//here is the return statement and it data.rows affected displays the number of rows available instead of having it display all of them in which it will look ugly
	//if you want to see change it from data.RowsAffected to all_members
	c.JSON(http.StatusOK, data.RowsAffected)
	
}

func Get_one_users(c *gin.Context) {
	db, err := middlewares.DbConnection()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	var one_member models.Members

	get_with_id := c.Param("id")

	the_member := db.Find(&one_member, get_with_id)

	if the_member.Error != nil{
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	
	c.JSON(http.StatusOK, one_member)
	
}
