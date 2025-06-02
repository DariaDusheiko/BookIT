package main

// тут для тебя просто необходимо раскомментить роуты удаления и получения информации

import (
	"log"

	"github.com/BookIT/backend/config"
	"github.com/gin-contrib/cors"

	"github.com/BookIT/backend/internal/app/handlers/bookings"
	"github.com/BookIT/backend/internal/app/handlers/tables"
	"github.com/BookIT/backend/internal/app/handlers/users"

	repositories "github.com/BookIT/backend/internal/app/repository"
	"github.com/BookIT/backend/internal/app/services"

	"github.com/BookIT/backend/internal/pkg/db"
	"github.com/BookIT/backend/internal/pkg/middleware"

	"github.com/gin-gonic/gin"

	_ "github.com/BookIT/backend/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Config error: %v", err)
	}

	if err := db.Init(cfg.DB.DSN()); err != nil {
		log.Fatalf("DB init error: %v", err)
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"X-Auth-Token", "Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"*"},
	}))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	setupRoutes(r)

	log.Printf("Server starting on %s:%s", cfg.App.Host, cfg.App.Port)
	if err := r.Run(cfg.App.Host + ":" + cfg.App.Port); err != nil {
		log.Fatal(err)
	}
}

func setupRoutes(r *gin.Engine) {
	userRepo := repositories.NewUserRepository(db.DB)
	bookingRepo := repositories.NewBookingRepository(db.DB)
	tableRepo := repositories.NewTableRepository(db.DB)

	userService := services.NewUserService(userRepo)
	bookingService := services.NewBookingService(bookingRepo, tableRepo)
	tableService := services.NewTableService(tableRepo, bookingRepo)

	userHandler := users.NewUserHandler(userService)
	bookingHandler := bookings.NewBookingHandler(bookingService)
	tableHandler := tables.NewTableHandler(tableService)

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/token", userHandler.Authenticate)
	}

	bookingGroup := r.Group("/booking")
	bookingGroup.Use(middleware.AuthMiddleware())
	{
		bookingGroup.POST("/", bookingHandler.CreateBooking)
		bookingGroup.DELETE("/", bookingHandler.DeleteBooking)
		bookingGroup.POST("/info", bookingHandler.GetUserBookings)
	}

	tableGroup := r.Group("/tables")
	tableGroup.Use(middleware.AuthMiddleware())
	{
		tableGroup.POST("/", tableHandler.GetTables)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
}
