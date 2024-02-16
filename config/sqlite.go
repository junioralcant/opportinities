package config

import (
	"github.com/junioralcant/opportinities.git/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InicializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")

	dns := "host=localhost user=postgres password=toor dbname=opportunities port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

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
