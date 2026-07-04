package certificate_handler

import (
	"strconv"

	"portfolio-api/base/helpers/error_helper"
	"portfolio-api/base/helpers/http_helper"
	"portfolio-api/constants"
	"portfolio-api/modules/certificate/certificate_dto"
	"portfolio-api/modules/certificate/certificate_service"

	"github.com/gin-gonic/gin"
)

// CertificateHandler exposes developer-scoped certificate endpoints.
type CertificateHandler interface {
	GetByUser(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type certificateHandlerImpl struct {
	service certificate_service.CertificateService
}

// NewCertificateHandler builds a CertificateHandler.
func NewCertificateHandler(service certificate_service.CertificateService) CertificateHandler {
	return &certificateHandlerImpl{service: service}
}

func (h *certificateHandlerImpl) GetByUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("invalid developer id"))
		return
	}
	result, err := h.service.GetByUser(userID)
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "OK", result)
}

func (h *certificateHandlerImpl) Create(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("invalid developer id"))
		return
	}
	var request certificate_dto.CertificateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("name and issuer are required"))
		return
	}
	result, err := h.service.Create(userID, request)
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_CREATED, "Certificate created", result)
}

func (h *certificateHandlerImpl) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("invalid certificate id"))
		return
	}
	var request certificate_dto.CertificateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("name and issuer are required"))
		return
	}
	result, err := h.service.Update(id, request)
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "Certificate updated", result)
}

func (h *certificateHandlerImpl) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("invalid certificate id"))
		return
	}
	if err := h.service.Delete(id); err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "Certificate deleted", nil)
}
