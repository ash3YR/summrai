package main

import (
	"net/http"

	"summrai-backend/internal/config"
	"summrai-backend/internal/db"

	"summrai-backend/internal/handlers"
	"summrai-backend/internal/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	database := db.NewDB(cfg)                            // ye add kra baadme

	userRepo := repository.NewUserRepository(database)   // ye add kra baadme


	userHandler := handlers.NewUserHandler(userRepo)      // ye add kra baadme



	defer database.Close()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "backend working",
		})
	})

	r.GET("/health", func(c *gin.Context) {
		err := database.Ping(c)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "database down",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "database healthy",
		})
	})


	r.POST("/users", userHandler.CreateUser)    //ye add kra baadme


	r.Run(":" + cfg.Port)


}