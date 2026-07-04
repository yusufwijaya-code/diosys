package service_handler

import (
	"strconv"

	"portfolio-api/base/helpers/error_helper"
	"portfolio-api/base/helpers/http_helper"
	"portfolio-api/constants"
	"portfolio-api/modules/service/service_dto"
	"portfolio-api/modules/service/service_service"

	"github.com/gin-gonic/gin"
)

// ServiceHandler exposes agency service endpoints.
type ServiceHandler interface {
	GetPublic(c *gin.Context)
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type serviceHandlerImpl struct {
	service service_service.Service
}

// NewServiceHandler builds a ServiceHandler.
func NewServiceHandler(service service_service.Service) ServiceHandler {
	return &serviceHandlerImpl{service: service}
}

func (h *serviceHandlerImpl) GetPublic(c *gin.Context) {
	result, err := h.service.GetPublic()
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "OK", result)
}

func (h *serviceHandlerImpl) GetAll(c *gin.Context) {
	result, err := h.service.GetAll()
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "OK", result)
}

func (h *serviceHandlerImpl) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("invalid service id"))
		return
	}
	result, err := h.service.GetByID(id)
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "OK", result)
}

func (h *serviceHandlerImpl) Create(c *gin.Context) {
	var request service_dto.ServiceRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("title is required"))
		return
	}
	result, err := h.service.Create(request)
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_CREATED, "Service created", result)
}

func (h *serviceHandlerImpl) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("invalid service id"))
		return
	}
	var request service_dto.ServiceRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("title is required"))
		return
	}
	result, err := h.service.Update(id, request)
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "Service updated", result)
}

func (h *serviceHandlerImpl) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("invalid service id"))
		return
	}
	if err := h.service.Delete(id); err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "Service deleted", nil)
}
