package controllers

import (
	"api/initializers"
	leaderboard "leaderboard/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTopN(c *gin.Context) {
	var topInt int = 10
	top := c.Query("top")
	if top != "" {
		topIntParsed, err := strconv.Atoi(top)
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": "Invalid top value"})
		}
		topInt = topIntParsed
	}
	results, err := leaderboard.GetTopNLeaderboard(initializers.RedisClient, topInt)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Invalid top value"})
	}

	c.JSON(http.StatusOK, results)
}

func GetUserRankScore(c *gin.Context) {
	userId := c.Param("id")

	rankScore, err := leaderboard.GetUserRankAndScore(initializers.RedisClient, userId)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Invalid top value"})
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
		c.AbortWithStatusJSON(500, gin.H{"error": "Invalid top value"})
	}

	c.JSON(http.StatusOK, rankScore)
}
