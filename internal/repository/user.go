package repository

import (
	"hotel-booking/internal/db"
	"hotel-booking/internal/errs"
	"hotel-booking/internal/models"
)

func GetUserByUsername(username string) (models.User, error) {
	var user models.User
	err := db.GetDBConn().Get(&user, `SELECT * FROM users WHERE username=$1 AND deleted_at IS NULL`, username)
	if err != nil {
		return models.User{}, errs.ErrNotFound
	}
	return user, nil
}

func CreateUser(user models.User) error {
	_, err := db.GetDBConn().Exec(
		`INSERT INTO users (full_name, username, password) VALUES ($1, $2, $3)`,
		user.FullName, user.Username, user.Password)
	return err
}

func GetUserByUsernameAndPassword(username, password string) (models.User, error) {
	var user models.User
	err := db.GetDBConn().Get(&user,
		`SELECT * FROM users WHERE username=$1 AND password=$2 AND deleted_at IS NULL`, username, password)
	if err != nil {
		return models.User{}, errs.ErrNotFound
	}
	return user, nil
}

func GetUserByID(userID int) (models.User, error) {
	var u models.User
	err := db.GetDBConn().Get(&u,
		`SELECT id, full_name, username, created_at FROM users WHERE id=$1 AND deleted_at IS NULL`, userID)
	if err != nil {
		return models.User{}, errs.ErrNotFound
	}
	return u, nil
}
