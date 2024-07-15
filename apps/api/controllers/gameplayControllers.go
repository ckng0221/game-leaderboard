package controllers

import (
	"api/initializers"
	"api/models"
	"api/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	leaderboard "leaderboard/utils"

	"github.com/gin-gonic/gin"
)

func GetAllGameplays(c *gin.Context) {
	userId := c.Query("user_id")

	m := make(map[string]interface{})

	if userId != "" {
		m["user_id"] = userId
	}

	var gameplays []models.Gameplay
	initializers.Db.Scopes(utils.Paginate(c)).Where(m).Find(&gameplays)

	c.JSON(http.StatusOK, gameplays)
}

func CreateGameplay(c *gin.Context) {
	var gameplay models.Gameplay

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	err = json.Unmarshal(body, &gameplay)

	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	// Fixed 10 score for 1 game
	gameplay.Score = 10

	result := initializers.Db.Create(&gameplay)
	if result.Error != nil {
		c.AbortWithStatus(500)
		return
	}
	go leaderboard.IncrementUserScore(initializers.RedisClient, gameplay.UserID, gameplay.Score)

	c.JSON(http.StatusCreated, gameplay)
}

func GetOneGameplay(c *gin.Context) {
	id := c.Param("id")

	var gameplay models.Gameplay
	result := initializers.Db.First(&gameplay, id)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	if gameplay.ID == 0 {
		c.JSON(http.StatusOK, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gameplay)
}

// func UpdateOneGameplay(c *gin.Context) {
// 	// get the id
// 	id := c.Param("id")
// 	var gameplay models.Gameplay
// 	initializers.Db.First(&gameplay, id)

// 	body, err := io.ReadAll(c.Request.Body)
// 	if err != nil {
// 		c.AbortWithError(400, err)
// 		return
// 	}
// 	var gameplayM map[string]interface{}

// 	err = json.Unmarshal(body, &gameplayM)

// 	if err != nil {
// 		c.AbortWithError(400, err)
// 		return
// 	}

// 	initializers.Db.Model(&gameplay).Updates(&gameplayM)

// 	c.JSON(200, gameplay)
// }

// func DeleteOneGameplay(c *gin.Context) {
// 	id := c.Param("id")

// 	initializers.Db.Delete(&models.Gameplay{}, id)

// 	// response
// 	c.Status(202)
// }
