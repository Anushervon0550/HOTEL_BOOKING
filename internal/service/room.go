package service

import (
	"hotel-booking/internal/models"
	"hotel-booking/internal/repository"
)

func GetAllRooms() ([]models.Room, error) {
	return repository.GetAllRooms()
}

func GetRoomByID(id int) (models.Room, error) {
	return repository.GetRoomByID(id)
}
