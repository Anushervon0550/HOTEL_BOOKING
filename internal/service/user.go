package service

import (
	"errors"
	"fmt" // Для отладочного вывода
	"hotel-booking/internal/errs"
	"hotel-booking/internal/models"
	"hotel-booking/internal/repository"
	"hotel-booking/internal/utils"
)

func CreateUser(u models.User) error {
	// Проверка, существует ли пользователь
	_, err := repository.GetUserByUsername(u.Username)
	if err == nil {
		return errs.ErrUserAlreadyExists
	}
	if !errors.Is(err, errs.ErrNotFound) {
		return err
	}

	// Отладка — вывод пароля до хеширования
	fmt.Printf("DEBUG: Пароль до хеширования: %s\n", u.Password)

	// Хеширование пароля
	u.Password = utils.GenerateHash(u.Password)

	// Отладка — вывод хеша пароля
	fmt.Printf("DEBUG: Хеш пароля для сохранения: %s\n", u.Password)

	// Сохраняем пользователя
	return repository.CreateUser(u)
}

func GetUserByUsernameAndPassword(username, password string) (models.User, error) {
	// Отладка — вывод пароля до хеширования
	fmt.Printf("DEBUG: Пароль для входа (до хеширования): %s\n", password)

	// Хеширование пароля
	password = utils.GenerateHash(password)

	// Отладка — вывод хеша пароля
	fmt.Printf("DEBUG: Пароль для входа (после хеширования): %s\n", password)

	// Проверка в базе
	user, err := repository.GetUserByUsernameAndPassword(username, password)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return models.User{}, errs.ErrIncorrectUsernameOrPassword
		}
		return models.User{}, err
	}
	return user, nil
}

func GetUserByID(userID int) (models.User, error) {
	return repository.GetUserByID(userID)
}
