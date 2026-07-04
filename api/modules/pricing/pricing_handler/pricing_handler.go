package pricing_handler

import (
	"strconv"

	"portfolio-api/base/helpers/error_helper"
	"portfolio-api/base/helpers/http_helper"
	"portfolio-api/constants"
	"portfolio-api/modules/pricing/pricing_dto"
	"portfolio-api/modules/pricing/pricing_service"

	"github.com/gin-gonic/gin"
)

// PricingHandler exposes pricing plan endpoints.
type PricingHandler interface {
	GetPublic(c *gin.Context)
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type pricingHandlerImpl struct {
	service pricing_service.PricingService
}

// NewPricingHandler builds a PricingHandler.
func NewPricingHandler(service pricing_service.PricingService) PricingHandler {
	return &pricingHandlerImpl{service: service}
}

func (h *pricingHandlerImpl) GetPublic(c *gin.Context) {
	result, err := h.service.GetPublic()
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "OK", result)
}

func (h *pricingHandlerImpl) GetAll(c *gin.Context) {
	result, err := h.service.GetAll()
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "OK", result)
}

func (h *pricingHandlerImpl) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("invalid plan id"))
		return
	}
	result, err := h.service.GetByID(id)
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "OK", result)
}

func (h *pricingHandlerImpl) Create(c *gin.Context) {
	var request pricing_dto.PricePlanRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("title is required"))
		return
	}
	result, err := h.service.Create(request)
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_CREATED, "Pricing plan created", result)
}

func (h *pricingHandlerImpl) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("invalid plan id"))
		return
	}
	var request pricing_dto.PricePlanRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("title is required"))
		return
	}
	result, err := h.service.Update(id, request)
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "Pricing plan updated", result)
}

func (h *pricingHandlerImpl) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("invalid plan id"))
		return
	}
	if err := h.service.Delete(id); err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "Pricing plan deleted", nil)
}
