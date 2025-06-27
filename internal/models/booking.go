package models

import "time"

type Booking struct {
	ID        int        `json:"id" db:"id"`
	UserID    int        `json:"user_id" db:"user_id"`
	RoomID    int        `json:"room_id" db:"room_id"`
	StartDate time.Time  `json:"start_date" db:"start_date"`
	EndDate   time.Time  `json:"end_date" db:"end_date"`
	Status    string     `json:"status" db:"status"` // например "booked", "cancelled"
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"-" db:"deleted_at"`
}

// BookingRequest используется для создания бронирования через API
type BookingRequest struct {
	RoomID    int    `json:"room_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}
