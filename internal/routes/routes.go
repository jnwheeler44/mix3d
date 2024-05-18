package routes

import (
	"mix3d/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes() *gin.Engine {
	router := gin.Default()

	// Person routes
	router.POST("/api/people", controllers.CreateUser)
	router.GET("/api/people", controllers.GetPeople)
	router.GET("/api/people/:id", controllers.GetPerson)

	// Mixer routes
	router.POST("/api/mixers", controllers.CreateMixer)
	router.GET("/api/mixers", controllers.GetMixers)
	router.GET("/api/mixers/:id", controllers.GetMixer)

	return router
}
