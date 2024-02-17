package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/junioralcant/opportinities.git/domain/contracts"
)

type ListOpeningController struct {
	ListOpeningUseCase contracts.IListOpening
}

func (l *ListOpeningController) Handle(ctx *gin.Context) {
	openings := l.ListOpeningUseCase.List()
	sendSuccess(ctx, "list-openings", openings)
}
