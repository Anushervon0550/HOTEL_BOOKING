package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"hotel-booking/internal/configs"
	"hotel-booking/logger"
	"os"
)

var db *sqlx.DB

func ConnectDB() error {
	cfg := configs.AppSettings.PostgresParams

	// Получаем пароль из переменной окружения DB_PASSWORD
	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		logger.Error.Println(" DB_PASSWORD не найден в переменных окружения")
	}

	// Формируем DSN строку для подключения к Postgres
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, password, cfg.Database)

	logger.Info.Printf(" DSN для подключения к БД: %s", dsn)

	var err error
	db, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		logger.Error.Printf(" Error connecting to DB: %v", err)
		return err
	}

	// Проверяем подключение к БД
	if err = db.Ping(); err != nil {
		logger.Error.Printf(" Error pinging DB: %v", err)
		return err
	}

	logger.Info.Println(" Успешное подключение к базе данных")
	return nil
}

func GetDBConn() *sqlx.DB {
	return db
}

func CloseDB() error {
	if db != nil {
		return db.Close()
	}
	return nil
}
