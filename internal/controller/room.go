package controller

import (
	"github.com/gin-gonic/gin"
	"hotel-booking/internal/models"
	"hotel-booking/internal/service"
	"net/http"
	"strconv"
)

// GetAllRooms godoc
// @Summary Получить список всех комнат
// @Description Возвращает список всех комнат
// @Tags rooms
// @Produce json
// @Success 200 {array} models.Room
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Router /rooms [get]
func GetAllRooms(c *gin.Context) {
	rooms, err := service.GetAllRooms()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rooms)
}

// GetRoomByID godoc
// @Summary Получить комнату по ID
// @Description Возвращает информацию о комнате по её ID
// @Tags rooms
// @Produce json
// @Param id path int true "ID комн��ты"
// @Success 200 {object} models.Room
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Security ApiKeyAuth
// @Router /rooms/{id} [get]
func GetRoomByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid room id"})
		return
	}

	room, err := service.GetRoomByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "room not found"})
		return
	}
	c.JSON(http.StatusOK, room)
}

// CreateRoom godoc
// @Summary Создать комнату
// @Description Создает новую комнату
// @Tags rooms
// @Accept json
// @Produce json
// @Param room body models.Room true "Данные комнаты"
// @Success 201 {object} models.Room
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Router /rooms [post]
func CreateRoom(c *gin.Context) {
	var room models.Room
	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	created, err := service.CreateRoom(room)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)
}

// UpdateRoom godoc
// @Summary Обновить комнату
// @Description Обновляет данные комнаты по ID
// @Tags rooms
// @Accept json
// @Produce json
// @Param id path int true "ID комнаты"
// @Param room body models.Room true "Данные комнаты"
// @Success 200 {object} models.Room
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Router /rooms/{id} [put]
func UpdateRoom(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid room id"})
		return
	}
	var room models.Room
	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	room.ID = id
	updated, err := service.UpdateRoom(room)
	if err != nil {
		if err.Error() == "not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "room not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}

// DeleteRoom godoc
// @Summary Удалить комнату
// @Description Удаляет комнату по ID
// @Tags rooms
// @Produce json
// @Param id path int true "ID комнаты"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Router /rooms/{id} [delete]
func DeleteRoom(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid room id"})
		return
	}
	err = service.DeleteRoom(id)
	if err != nil {
		if err.Error() == "not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "room not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "room deleted"})
}
