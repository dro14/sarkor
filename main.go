package main

import (
	"log"

	"github.com/dro14/sarkor/database"
	"github.com/dro14/sarkor/handlers"
	"github.com/dro14/sarkor/middleware"
	"github.com/gin-gonic/gin"
)

func main() {

	database.Init()
	r := gin.Default()

	// public routes
	r.POST("/user/register", handlers.RegisterUser)
	r.POST("/user/auth", handlers.AuthenticateUser)

	// authorized routes
	authorized := r.Group("/user")
	authorized.Use(middleware.Auth)
	{
		authorized.GET("/:name", handlers.GetUserByName)
		authorized.POST("/phone", handlers.AddPhone)
		authorized.GET("/phone", handlers.GetPhone)
		authorized.PUT("/phone", handlers.UpdatePhone)
		authorized.DELETE("/phone/:phone_id", handlers.DeletePhone)
	}

	err := r.Run(":8080")
	if err != nil {
		log.Fatal("can't run server:", err)
	}
}
