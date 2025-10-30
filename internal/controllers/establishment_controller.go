// controllers/Establishment_controller.go
package controllers

import (
	"fmt"
	"net/http"
	"socialVoleiAPI/internal/models"
	"socialVoleiAPI/internal/services"

	"github.com/gin-gonic/gin"
)

type EstablishmentController struct {
	service *services.EstablishmentService
}

func NewEstablishmentController(s *services.EstablishmentService) *EstablishmentController {
	return &EstablishmentController{service: s}
}

func (uc *EstablishmentController) CreateEstablishment(ctx *gin.Context) {

	var establishment models.Establishment

	if err := ctx.ShouldBindJSON(&establishment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uc.service.CreateEstablishment(&establishment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, EstablishmentToResponse(&establishment))
}

func (uc *EstablishmentController) BulkCreateEstablishment(ctx *gin.Context) {
	var establishments []models.Establishment

	if err := ctx.ShouldBindJSON(&establishments); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uc.service.BulkCreateEstablishment(establishments); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func (uc *EstablishmentController) GetEstablishments(ctx *gin.Context) {

	estabs, err := uc.service.GetAllEstablishments()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error: ": err.Error()})
		return
	}

	establishmentsResponses := make([]models.EstablishmentResponse, len(estabs))

	for idx, Establishment := range estabs {
		establishmentsResponses[idx] = EstablishmentToResponse(&Establishment)
	}

	ctx.JSON(http.StatusAccepted, establishmentsResponses)
}

func (uc *EstablishmentController) GetEstablishmentByID(ctx *gin.Context) {

	establishment, err := uc.service.GetEstablishmentByID(ctx.Param("id"))

	if err != nil {
		fmt.Println(err.Error())
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error: ": "Id Não Encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error: ": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, EstablishmentToResponse(&establishment))
}

func EstablishmentToResponse(estab *models.Establishment) models.EstablishmentResponse {
	return models.EstablishmentResponse{
		ID:        estab.ID,
		Name:      estab.Name,
		Email:     estab.Email,
		Phone:     estab.Phone,
		CNPJ:      estab.CNPJ,
		Latitude:  estab.Latitude,
		Longitude: estab.Longitude,
	}
}
