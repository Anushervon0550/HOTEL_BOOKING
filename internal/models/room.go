package models

import "time"

type Room struct {
	ID         int        `json:"id" db:"id"`
	RoomNumber string     `json:"room_number" db:"room_number"`
	Type       string     `json:"type" db:"type"`
	Price      float64    `json:"price" db:"price"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt  *time.Time `json:"-" db:"deleted_at"`
}
