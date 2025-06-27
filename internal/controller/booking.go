package controller

import (
	"github.com/gin-gonic/gin"
	"hotel-booking/internal/models"
	"hotel-booking/internal/service"
	"net/http"
	"strconv"
	"time"
)

// CreateBooking godoc
// @Summary Создать бронирование
// @Description Создает новое бронирование для пользователя
// @Tags bookings
// @Accept json
// @Produce json
// @Param booking body models.BookingRequest true "Данные для бронирования"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Router /bookings [post]
func CreateBooking(c *gin.Context) {
	var input models.BookingRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	userID := c.GetInt("userID")

	startDate, err := time.Parse("2006-01-02", input.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid start_date format, use YYYY-MM-DD"})
		return
	}

	endDate, err := time.Parse("2006-01-02", input.EndDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid end_date format, use YYYY-MM-DD"})
		return
	}

	if !endDate.After(startDate) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "end_date must be after start_date"})
		return
	}

	booking := models.Booking{
		UserID:    userID,
		RoomID:    input.RoomID,
		StartDate: startDate,
		EndDate:   endDate,
	}

	if err := service.CreateBooking(booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "booking created"})
}

// GetMyBookings godoc
// @Summary Мои бронирования
// @Description Получить список всех бронирований пользователя
// @Tags bookings
// @Produce json
// @Success 200 {array} models.Booking
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Router /bookings [get]
func GetMyBookings(c *gin.Context) {
	userID := c.GetInt("userID")

	bookings, err := service.GetBookingsByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bookings)
}

// CancelBooking godoc
// @Summary Отмена бронирования
// @Description Отменяет бронирование пользователя по ID
// @Tags bookings
// @Produce json
// @Param id path int true "ID бронирования"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Router /bookings/{id} [delete]
func CancelBooking(c *gin.Context) {
	userID := c.GetInt("userID")

	bookingIDStr := c.Param("id")
	bookingID, err := strconv.Atoi(bookingIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid booking id"})
		return
	}

	if err := service.CancelBooking(bookingID, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "booking cancelled"})
}

// GetAllBookings godoc
// @Summary Получить все бронирования
// @Description Доступно только для manager и admin
// @Tags admin
// @Produce json
// @Success 200 {array} models.Booking
// @Failure 403 {object} map[string]string
// @Security ApiKeyAuth
// @Router /bookings/all [get]
func GetAllBookings(c *gin.Context) {
	bookings, err := service.GetAllBookings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bookings)
}
