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

func CreateRoom(room models.Room) (models.Room, error) {
	return repository.CreateRoom(room)
}

func UpdateRoom(room models.Room) (models.Room, error) {
	return repository.UpdateRoom(room)
}

func DeleteRoom(id int) error {
	return repository.DeleteRoom(id)
}
