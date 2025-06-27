package controller

import (
	"github.com/gin-gonic/gin"
	"hotel-booking/internal/service"
	"net/http"
)

// GetMyProfile godoc
// @Summary Получить профиль пользователя
// @Description Возвращает информацию о текущем пользователе
// @Tags user
// @Produce json
// @Success 200 {object} models.User
// @Failure 404 {object} map[string]string
// @Security ApiKeyAuth
// @Router /profile [get]
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
