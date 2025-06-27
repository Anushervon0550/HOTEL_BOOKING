package controller

import (
	"github.com/gin-gonic/gin"
	"hotel-booking/internal/service"
	"net/http"
	"strconv"
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

// GetAllUsers godoc
// @Summary Получить всех пользователей
// @Description Доступно только для admin
// @Tags admin
// @Produce json
// @Success 200 {array} models.User
// @Failure 403 {object} map[string]string
// @Security ApiKeyAuth
// @Router /users [get]
func GetAllUsers(c *gin.Context) {
	users, err := service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// DeleteUser godoc
// @Summary Удалить пользователя
// @Description Доступно только для admin
// @Tags admin
// @Produce json
// @Param id path int true "ID пользователя"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Security ApiKeyAuth
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	if err := service.DeleteUser(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}

// UpdateUserRole godoc
// @Summary Изменить роль пользователя
// @Description Доступно только для admin
// @Tags admin
// @Accept json
// @Produce json
// @Param id path int true "ID пользователя"
// @Param role body struct{Role string `json:"role"`} true "Новая роль"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Security ApiKeyAuth
// @Router /users/{id}/role [put]
func UpdateUserRole(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	var req struct {
		Role string `json:"role"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.Role == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid role"})
		return
	}
	if err := service.UpdateUserRole(id, req.Role); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "role updated"})
}
