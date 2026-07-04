package education_handler

import (
	"strconv"

	"portfolio-api/base/helpers/error_helper"
	"portfolio-api/base/helpers/http_helper"
	"portfolio-api/constants"
	"portfolio-api/modules/education/education_dto"
	"portfolio-api/modules/education/education_service"

	"github.com/gin-gonic/gin"
)

// EducationHandler exposes developer-scoped education endpoints.
type EducationHandler interface {
	GetByUser(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type educationHandlerImpl struct {
	service education_service.EducationService
}

// NewEducationHandler builds an EducationHandler.
func NewEducationHandler(service education_service.EducationService) EducationHandler {
	return &educationHandlerImpl{service: service}
}

func (h *educationHandlerImpl) GetByUser(c *gin.Context) {
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

func (h *educationHandlerImpl) Create(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("invalid developer id"))
		return
	}
	var request education_dto.EducationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("degree and institution are required"))
		return
	}
	result, err := h.service.Create(userID, request)
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_CREATED, "Education created", result)
}

func (h *educationHandlerImpl) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("invalid education id"))
		return
	}
	var request education_dto.EducationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("degree and institution are required"))
		return
	}
	result, err := h.service.Update(id, request)
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "Education updated", result)
}

func (h *educationHandlerImpl) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("invalid education id"))
		return
	}
	if err := h.service.Delete(id); err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "Education deleted", nil)
}
