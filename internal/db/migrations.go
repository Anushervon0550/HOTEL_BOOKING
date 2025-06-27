package db

import (
	"hotel-booking/logger"
)

func InitMigrations() error {
	usersTable := `
CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  full_name VARCHAR(255) NOT NULL,
  username VARCHAR(255) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  role VARCHAR(50) NOT NULL DEFAULT 'user',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP
);
`
	if _, err := db.Exec(usersTable); err != nil {
		logger.Error.Printf("Error creating users table: %v", err)
		return err
	}

	roomsTable := `
CREATE TABLE IF NOT EXISTS rooms (
  id SERIAL PRIMARY KEY,
  room_number VARCHAR(50) NOT NULL UNIQUE,
  type VARCHAR(100) NOT NULL,
  price NUMERIC(10,2) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP
);
`
	if _, err := db.Exec(roomsTable); err != nil {
		logger.Error.Printf("Error creating rooms table: %v", err)
		return err
	}

	bookingsTable := `
CREATE TABLE IF NOT EXISTS bookings (
  id SERIAL PRIMARY KEY,
  user_id INT REFERENCES users(id),
  room_id INT REFERENCES rooms(id),
  start_date DATE NOT NULL,
  end_date DATE NOT NULL,
  status VARCHAR(50) NOT NULL DEFAULT 'booked',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP
);
`
	if _, err := db.Exec(bookingsTable); err != nil {
		logger.Error.Printf("Error creating bookings table: %v", err)
		return err
	}

	return nil
}
