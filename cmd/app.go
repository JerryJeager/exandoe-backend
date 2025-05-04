package cmd

import (
	"log"
	"os"

	"github.com/JerryJeager/exandoe-backend/manualwire"
	"github.com/JerryJeager/exandoe-backend/middleware"
	"github.com/gin-gonic/gin"
)

func ExecuteApiRoutes() {
	router := gin.Default()

	router.Use(middleware.CORSMiddleware())

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome to exandoe",
		})
	})
	
	userController := manualwire.GetUserController()
	gameController := manualwire.GetGameController()
	
	api := router.Group("/api/v1")
	users := api.Group("/users")
	games := api.Group("/games")


	users.GET("/lobby", userController.Signin)
		
	games.GET("/play", gameController.Play)
	

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
