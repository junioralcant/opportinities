package contracts

import "github.com/junioralcant/opportinities.git/domain/models"

type IListOpening interface {
	List() []models.Opening
}
