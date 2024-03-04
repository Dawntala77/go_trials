package handlers

import (
	"encoding/json"
	"example.com/myproject/database"
	"example.com/myproject/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func SignUp(c *gin.Context) {
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

	newMembers := models.Members{
		Name:         member.Name,
		EmailAddress: member.EmailAddress,
		Password:     member.Password,
	}

	db, err := database.ConnectT()
	if err != nil {
		fmt.Println("DB Connection failed")
		panic(err)
	}

	if err = db.Create(&newMembers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "User created successfully")
}

func GetUsers(c *gin.Context) {
	db, err := database.ConnectT()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	var allMembers []models.Members
	data := db.Find(&allMembers)

	if data.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	//here is the return statement and it data.rows affected displays the number of rows available instead of having it display all of them in which it will look ugly
	//if you want to see change it from data.RowsAffected to all_members
	c.JSON(http.StatusOK, allMembers)

}

func GetUser(c *gin.Context) {
	db, err := database.ConnectT()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	var user models.Members

	userId := c.Param("id")

	member := db.Find(&user, userId)
	if member.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	db, err := database.ConnectT()
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

func DeleteUser(c *gin.Context) {
	db, err := database.ConnectT()
	if err != nil {
		panic(err)
	}

	member_delete := &models.Members{}
	id := c.Param("id")

	if err := db.Unscoped().Delete(member_delete, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "record is errased in the database")

}
