package routes

import (
	"api/controllers"

	"github.com/gin-gonic/gin"
)

func GameplayRoutes(r *gin.Engine) *gin.Engine {

	r.GET("/gameplays", controllers.GetAllGameplays)
	r.POST("/gameplays", controllers.CreateGameplay)
	r.GET("/gameplays/:id", controllers.GetOneGameplay)
	// r.PATCH("/gameplays/:id", controllers.UpdateOneGameplay)
	// r.DELETE("/gameplays/:id", controllers.DeleteOneGameplay)

	// Auth

	return r
}
