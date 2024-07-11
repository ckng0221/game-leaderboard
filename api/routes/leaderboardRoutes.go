package routes

import (
	"api/controllers"

	"github.com/gin-gonic/gin"
)

func LeaderboardRoutes(r *gin.Engine) *gin.Engine {

	r.GET("/leaderboard", controllers.GetTopN)
	r.GET("/leaderboard/:rank", controllers.GetTopN)

	return r
}
