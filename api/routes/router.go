package routes

import (
	"api/middleware"

	"github.com/gin-gonic/gin"
)

// @Summary Health check
// @Tags Default
// @Produce json
// @Success 200
// @Router / [get]
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello Game Leaderboard")
	})

	UserRoutes(r)
	GameplayRoutes(r)
	LeaderboardRoutes(r)

	return r
}
