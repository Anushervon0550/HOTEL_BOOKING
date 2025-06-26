package controller

import (
	"github.com/gin-gonic/gin"
	"hotel-booking/internal/service"
	"net/http"
)

// GetMyProfile возвращает информацию о текущем пользователе
func GetMyProfile(c *gin.Context) {
	userID := c.GetInt("userID")

	user, err := service.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
