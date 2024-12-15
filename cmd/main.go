package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"

	"webfolio-backend/internal/domain"
	"webfolio-backend/internal/handler"
	"webfolio-backend/internal/infrastructures/database"
	"webfolio-backend/internal/repository"
	"webfolio-backend/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	database.ConnectDatabase()

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: database.DB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to initialize GORM with the database:", err)
	}

	if err := db.AutoMigrate(&domain.User{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.ExposeHeaders = []string{"Content-Length"}
	r.Use(cors.New(config))

	r.POST("/users", userHandler.CreateUser)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
