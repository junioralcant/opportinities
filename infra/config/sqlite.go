package config

import (
	"github.com/junioralcant/opportinities.git/domain/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InicializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")

	dsn := "host=localhost user=postgres password=toor dbname=opportunities port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.Errorf("sqlite initialization error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&models.Opening{})
	if err != nil {
		logger.Errorf("sqlite auto migration error: %v", err)
		return nil, err
	}

	return db, nil
}
