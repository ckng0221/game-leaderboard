package controllers

import (
	"api/initializers"
	"api/models"
	"api/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	email := c.Query("email")

	m := make(map[string]interface{})

	if email != "" {
		m["email"] = email
	}

	var users []models.User
	initializers.Db.Scopes(utils.Paginate(c)).Where(m).Find(&users)

	c.JSON(http.StatusOK, users)
}

// By admin only
func CreateUsers(c *gin.Context) {
	var users []models.User

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	err = json.Unmarshal(body, &users)

	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	result := initializers.Db.Create(&users)
	if result.Error != nil {
		c.AbortWithStatus(500)
		return
	}

	c.JSON(http.StatusCreated, users)
}

func GetOneUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	result := initializers.Db.First(&user, id)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	if user.ID == 0 {
		c.JSON(http.StatusOK, gin.H{})
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateOneUser(c *gin.Context) {
	// get the id
	id := c.Param("id")
	var user models.User
	initializers.Db.First(&user, id)

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}
	var userM map[string]interface{}

	err = json.Unmarshal(body, &userM)

	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	initializers.Db.Model(&user).Updates(&userM)

	c.JSON(200, user)
}

func DeleteOneUser(c *gin.Context) {
	id := c.Param("id")

	initializers.Db.Delete(&models.User{}, id)

	// response
	c.Status(202)
}

func GetUserRoles(c *gin.Context) {
	c.JSON(http.StatusOK, models.Roles)
}
