package repositories

import (
	"github.com/junioralcant/opportinities.git/domain/models"
)

type IOpeningRepository interface {
	ListOpening() []models.Opening
}
