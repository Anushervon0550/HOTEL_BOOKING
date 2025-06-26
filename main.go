package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hotel-booking/internal/configs"
	"hotel-booking/internal/controller"
	"hotel-booking/internal/db"
	"hotel-booking/logger"
)

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

	// Роуты аутентификации
	r.POST("/auth/sign-up", controller.SignUp)
	r.POST("/auth/sign-in", controller.SignIn)

	// Защищенные роуты
	auth := r.Group("/")
	auth.Use(controller.AuthMiddleware())

	// Роуты комнат
	auth.GET("/rooms", controller.GetAllRooms)
	auth.GET("/rooms/:id", controller.GetRoomByID)
	auth.GET("/profile", controller.GetMyProfile)

	// Роуты бронирования
	auth.POST("/bookings", controller.CreateBooking)
	auth.GET("/bookings", controller.GetMyBookings)
	auth.DELETE("/bookings/:id", controller.CancelBooking)

	port := configs.AppSettings.AppParams.PortRun
	logger.Info.Println("Starting server on port", port)
	if err := r.Run(port); err != nil {
		logger.Error.Println("Error starting server:", err)
	}
}
