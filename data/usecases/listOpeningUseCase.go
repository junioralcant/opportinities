package data

import (
	"github.com/junioralcant/opportinities.git/domain/models"
	"github.com/junioralcant/opportinities.git/infra/repositories"
)

type ListOpeningUseCase struct {
	OpeningRepository repositories.IOpeningRepository
}

func (l *ListOpeningUseCase) List() []models.Opening {
	return l.OpeningRepository.ListOpening()
}
