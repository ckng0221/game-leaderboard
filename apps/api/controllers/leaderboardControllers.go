package controllers

import (
	"api/initializers"
	"api/models"
	"fmt"
	leaderboard "leaderboard/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LeaderboardData struct {
	Rank     int    `json:"rank"`
	Username string `json:"username"`
	Score    int    `json:"score"`
}

func GetTopN(c *gin.Context) {
	var topInt int = 10
	top := c.Query("top")
	if top != "" {
		topIntParsed, err := strconv.Atoi(top)
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": "Invalid top value"})
			return
		}
		topInt = topIntParsed
	}
	results, err := leaderboard.GetTopNLeaderboard(initializers.RedisClient, topInt)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(500)
		return
	}

	// modify table to create leaderboard table
	var leaderboardTable []LeaderboardData
	var users []models.User
	var user_ids []string
	for _, result := range results {
		user_ids = append(user_ids, result.Member.(string))
	}
	err = initializers.Db.Find(&users, user_ids).Error
	if err != nil || len(users) != len(user_ids) {
		if len(users) != len(user_ids) {
			log.Println("Users and fetched users from db are not tally")
		} else {
			log.Println(err.Error())
		}
		c.AbortWithStatus(500)
		return
	}

	userid_hashtable := map[string]models.User{}
	for _, user := range users {
		userid_hashtable[fmt.Sprint(user.ID)] = user
	}

	for i, result := range results {
		leaderboard := LeaderboardData{
			Rank:     i + 1,
			Username: userid_hashtable[result.Member.(string)].Username,
			Score:    int(result.Score),
		}
		leaderboardTable = append(leaderboardTable, leaderboard)
	}

	c.JSON(http.StatusOK, leaderboardTable)
}

func GetUserRankScore(c *gin.Context) {
	userId := c.Param("id")

	rankScore, err := leaderboard.GetUserRankAndScore(initializers.RedisClient, userId)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(500)
		return
	}

	c.JSON(http.StatusOK, rankScore)
}

func GetUserByRank(c *gin.Context) {
	rank := c.Param("rank")
	rank_int, err := strconv.Atoi(rank)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid rank value"})
	}

	rankScore, err := leaderboard.GetUserByRank(initializers.RedisClient, rank_int)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(500)
		return
	}

	c.JSON(http.StatusOK, rankScore)
}
