package controllers

import (
	"github.com/junioralcant/opportinities.git/infra/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeController() {
	logger = config.GetLogger("controllers")
	db = config.GetSQLite()
}
