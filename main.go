package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "hotel-booking/docs" // импортируем сгенерированные swagger docs
	"hotel-booking/internal/configs"
	"hotel-booking/internal/controller"
	"hotel-booking/internal/db"
	"hotel-booking/logger"
)

// @title Hotel Booking API
// @version 1.0
// @description API для бронирования отеля
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	// Читаем конфигурацию
	if err := configs.ReadSettings(); err != nil {
		fmt.Println("Error reading configs:", err)
		return
	}

	// Инициализируем логгер
	logger.Init()

	// Подключаемся к базе данных
	if err := db.ConnectDB(); err != nil {
		logger.Error.Println("Error connecting to DB:", err)
		return
	}
	defer db.CloseDB()

	// Выполняем миграции
	if err := db.InitMigrations(); err != nil {
		logger.Error.Println("Error migrating DB:", err)
		return
	}

	// Устанавливаем режим Gin
	gin.SetMode(configs.AppSettings.AppParams.GinMode)

	r := gin.Default()

	// Swagger роут
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Роуты аутентификации
	r.POST("/auth/sign-up", controller.SignUp)
	r.POST("/auth/sign-in", controller.SignIn)

	// Защищенные роуты
	auth := r.Group("/")
	auth.Use(controller.AuthMiddleware())

	// Роуты комнат
	auth.GET("/rooms", controller.GetAllRooms)
	auth.GET("/rooms/:id", controller.GetRoomByID)
	// Только для manager и admin
	auth.POST("/rooms", controller.RequireRoles("admin", "manager"), controller.CreateRoom)
	auth.PUT("/rooms/:id", controller.RequireRoles("admin", "manager"), controller.UpdateRoom)
	auth.DELETE("/rooms/:id", controller.RequireRoles("admin", "manager"), controller.DeleteRoom)
	auth.GET("/profile", controller.GetMyProfile)

	// Роуты бронирования
	auth.POST("/bookings", controller.CreateBooking)
	auth.GET("/bookings", controller.GetMyBookings)
	auth.DELETE("/bookings/:id", controller.CancelBooking)
	auth.GET("/bookings/all", controller.RequireRoles("admin", "manager"), controller.GetAllBookings)

	// Роуты управления пользователями (только для admin)
	auth.GET("/users", controller.RequireRoles("admin"), controller.GetAllUsers)
	auth.DELETE("/users/:id", controller.RequireRoles("admin"), controller.DeleteUser)
	auth.PUT("/users/:id/role", controller.RequireRoles("admin"), controller.UpdateUserRole)

	port := configs.AppSettings.AppParams.PortRun
	logger.Info.Println("Starting server on port", port)
	if err := r.Run(port); err != nil {
		logger.Error.Println("Error starting server:", err)
	}
}
