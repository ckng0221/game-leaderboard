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
	userIds, err := leaderboard.GetTopNLeaderboard(initializers.RedisClient, topInt)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Invalid top value"})
	}

	c.JSON(http.StatusOK, userIds)
}
