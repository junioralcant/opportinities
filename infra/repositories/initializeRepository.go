package repositories

import (
	"github.com/junioralcant/opportinities.git/infra/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeRepository() {
	logger = config.GetLogger("controllers")
	db = config.GetSQLite()
}
