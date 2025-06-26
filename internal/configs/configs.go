package configs

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"hotel-booking/internal/models"
	"log"
	"os"
)

var AppSettings models.Configs

func ReadSettings() error {
	fmt.Println("Loading .env file")

	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env file not found, using system environment variables")
	}

	fmt.Println("Reading configs.json")
	configFile, err := os.Open("configs.json")
	if err != nil {
		return errors.New(fmt.Sprintf("Couldn't open config file: %s", err.Error()))
	}
	defer func() {
		if err := configFile.Close(); err != nil {
			log.Fatal("Couldn't close config file: ", err.Error())
		}
	}()

	if err = json.NewDecoder(configFile).Decode(&AppSettings); err != nil {
		return errors.New(fmt.Sprintf("Couldn't decode json config: %s", err.Error()))
	}

	return nil
}
