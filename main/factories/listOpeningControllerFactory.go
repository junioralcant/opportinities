package factories

import (
	data "github.com/junioralcant/opportinities.git/data/usecases"
	"github.com/junioralcant/opportinities.git/infra/repositories"
	"github.com/junioralcant/opportinities.git/presentation/controllers"
)

func ListOpeningControllerFactory() *controllers.ListOpeningController {
	repo := &repositories.OpeningRepository{}
	useCase := &data.ListOpeningUseCase{OpeningRepository: repo}
	listOpeningController := controllers.ListOpeningController{ListOpeningUseCase: useCase}
	return &listOpeningController
}
