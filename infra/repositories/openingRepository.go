package repositories

import "github.com/junioralcant/opportinities.git/domain/models"

type OpeningRepository struct {
}

func (o *OpeningRepository) ListOpening() []models.Opening {
	openings := []models.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		logger.Errorf("error in list openings: %v", err)
	}

	return openings
}
