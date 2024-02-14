package config

import (
	"os"

	"github.com/junioralcant/opportinities.git/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InicializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("database not found, creating...")

		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}

		file.Close()
	}

	db, err := gorm.Open(sqlite.Open("./db/main.db"), &gorm.Config{})

	if err != nil {
		logger.Errorf("sqlite initialization error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("sqlite auto migration error: %v", err)
		return nil, err
	}

	return db, nil
}