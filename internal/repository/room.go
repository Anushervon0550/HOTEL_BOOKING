package repository

import (
	"hotel-booking/internal/db"
	"hotel-booking/internal/errs"
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

func CreateRoom(room models.Room) (models.Room, error) {
	query := `INSERT INTO rooms (room_number, type, price) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at`
	var id int
	var createdAt, updatedAt string
	err := db.GetDBConn().QueryRowx(query, room.RoomNumber, room.Type, room.Price).Scan(&id, &createdAt, &updatedAt)
	if err != nil {
		return models.Room{}, err
	}
	room.ID = id
	// createdAt/updatedAt можно парсить в time.Time при необходимости
	return room, nil
}

func UpdateRoom(room models.Room) (models.Room, error) {
	query := `UPDATE rooms SET room_number=$1, type=$2, price=$3, updated_at=NOW() WHERE id=$4 AND deleted_at IS NULL RETURNING updated_at`
	var updatedAt string
	err := db.GetDBConn().QueryRowx(query, room.RoomNumber, room.Type, room.Price, room.ID).Scan(&updatedAt)
	if err != nil {
		return models.Room{}, err
	}
	return room, nil
}

func DeleteRoom(id int) error {
	query := `UPDATE rooms SET deleted_at=NOW() WHERE id=$1 AND deleted_at IS NULL`
	res, err := db.GetDBConn().Exec(query, id)
	if err != nil {
		return err
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return errs.ErrNotFound
	}
	return nil
}
