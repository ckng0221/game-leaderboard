package routes

import (
	"api/controllers"

	"github.com/gin-gonic/gin"
)

func LeaderboardRoutes(r *gin.Engine) *gin.Engine {

	r.GET("/leaderboard", controllers.GetTopN)
	r.GET("/leaderboard/users/:id", controllers.GetUserRankScore)
	r.GET("/leaderboard/ranks/:rank", controllers.GetUserByRank)

	return r
}
