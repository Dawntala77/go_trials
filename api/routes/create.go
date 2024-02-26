package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"example.com/myproject/api/middlewares"
	"example.com/myproject/api/models"

	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {

	jsonBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	var member models.Members

	err = json.Unmarshal(jsonBody, &member)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse JSON"})
		return
	}

	c.JSON(http.StatusOK, member)

	Newmembers := models.Members{Identity: member.Identity, Name: member.Name, Email_addres: member.Email_addres, Password: member.Password}

	db, err := middlewares.DbConnection()
	if err != nil {
		panic(err)
	}

	if err := db.Create(&Newmembers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "user created successfully")
}
