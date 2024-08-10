package controllers

import (
	"api/initializers"
	"api/models"
	"api/utils"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	gameplayScore := 10
	gameplay.Score = gameplayScore

	result := initializers.Db.Create(&gameplay)
	if result.Error != nil {
		log.Println(err)
		c.AbortWithStatus(500)
		return
	}

	// Publish score event to event queue
	err = initializers.RabbitMqObj.PublishScore(gameplay.UserID, gameplayScore)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(500)
		return
	}

	// go leaderboard.IncrementUserScore(initializers.RedisClient, gameplay.UserID, gameplay.Score)

	// Update score in DB
	var score models.Score
	month := time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.Now().UTC().Location())
	err = initializers.Db.Where("month = ? AND user_id = ?", month, gameplay.UserID).Find(&score).Error
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(500)
		return
	}
	if score.ID != 0 {
		// If existing, update score
		expression := "score + ?"
		err = initializers.Db.Model(&score).UpdateColumn("score", gorm.Expr(expression, gameplayScore)).Error
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(500)
			return
		}
	} else {
		score.UserID = gameplay.UserID
		score.Score = gameplayScore
		score.Month = &month

		err = initializers.Db.Create(&score).Error
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(500)
			return
		}
	}

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
