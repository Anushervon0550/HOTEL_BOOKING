package service

import (
	"errors"
	"hotel-booking/internal/models"
	"hotel-booking/internal/repository"
)

func CreateBooking(b models.Booking) error {
	// Проверка, что комната существует
	_, err := repository.GetRoomByID(b.RoomID)
	if err != nil {
		return err
	}

	// Проверка на конфликт дат
	conflict, err := repository.CheckBookingConflict(b.RoomID, b.StartDate, b.EndDate)
	if err != nil {
		return err
	}
	if conflict {
		return errors.New("room already booked for these dates")
	}

	b.Status = "booked"
	return repository.CreateBooking(b)
}

func GetBookingsByUserID(userID int) ([]models.Booking, error) {
	return repository.GetBookingsByUserID(userID)
}

func CancelBooking(bookingID, userID int) error {
	return repository.CancelBooking(bookingID, userID)
}
