package repository

import (
	"hotel-booking/internal/db"
	"hotel-booking/internal/models"
)

func GetAllRooms() ([]models.Room, error) {
	var rooms []models.Room
	err := db.GetDBConn().Select(&rooms, `SELECT * FROM rooms WHERE deleted_at IS NULL`)
	return rooms, err
}

func GetRoomByID(id int) (models.Room, error) {
	var room models.Room
	err := db.GetDBConn().Get(&room, `SELECT * FROM rooms WHERE id=$1 AND deleted_at IS NULL`, id)
	return room, err
}
