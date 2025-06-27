package repository

import (
	"hotel-booking/internal/db"
	"hotel-booking/internal/models"
	"time"
)

// Создать бронирование
func CreateBooking(b models.Booking) error {
	_, err := db.GetDBConn().Exec(`
INSERT INTO bookings (user_id, room_id, start_date, end_date, status) 
VALUES ($1, $2, $3, $4, $5)`,
		b.UserID, b.RoomID, b.StartDate, b.EndDate, b.Status)
	return err
}

// Получить бронирования по ID пользователя
func GetBookingsByUserID(userID int) ([]models.Booking, error) {
	var bookings []models.Booking
	err := db.GetDBConn().Select(&bookings,
		`SELECT * FROM bookings WHERE user_id=$1 AND deleted_at IS NULL ORDER BY start_date`, userID)
	return bookings, err
}

// Получить все бронирования (для manager/admin)
func GetAllBookings() ([]models.Booking, error) {
	var bookings []models.Booking
	err := db.GetDBConn().Select(&bookings, `SELECT * FROM bookings WHERE deleted_at IS NULL ORDER BY start_date`)
	return bookings, err
}

// Проверка конфликта бронирования
func CheckBookingConflict(roomID int, startDate, endDate time.Time) (bool, error) {
	var count int
	err := db.GetDBConn().Get(&count,
		`SELECT COUNT(*) FROM bookings 
     WHERE room_id=$1 AND deleted_at IS NULL AND status='booked' 
       AND NOT (end_date < $2 OR start_date > $3)`, roomID, startDate, endDate)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// Отмена бронирования
func CancelBooking(bookingID, userID int) error {
	res, err := db.GetDBConn().Exec(
		`UPDATE bookings SET status='cancelled', updated_at=CURRENT_TIMESTAMP 
     WHERE id=$1 AND user_id=$2 AND deleted_at IS NULL`, bookingID, userID)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return nil
	}
	return nil
}
