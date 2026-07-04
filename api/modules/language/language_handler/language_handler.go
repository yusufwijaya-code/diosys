package language_handler

import (
	"strconv"

	"portfolio-api/base/helpers/error_helper"
	"portfolio-api/base/helpers/http_helper"
	"portfolio-api/constants"
	"portfolio-api/modules/language/language_dto"
	"portfolio-api/modules/language/language_service"

	"github.com/gin-gonic/gin"
)

// LanguageHandler exposes developer-scoped language endpoints.
type LanguageHandler interface {
	GetByUser(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type languageHandlerImpl struct {
	service language_service.LanguageService
}

// NewLanguageHandler builds a LanguageHandler.
func NewLanguageHandler(service language_service.LanguageService) LanguageHandler {
	return &languageHandlerImpl{service: service}
}

func (h *languageHandlerImpl) GetByUser(c *gin.Context) {
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

func (h *languageHandlerImpl) Create(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("invalid developer id"))
		return
	}
	var request language_dto.LanguageRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("name and level are required"))
		return
	}
	result, err := h.service.Create(userID, request)
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_CREATED, "Language created", result)
}

func (h *languageHandlerImpl) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("invalid language id"))
		return
	}
	var request language_dto.LanguageRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("name and level are required"))
		return
	}
	result, err := h.service.Update(id, request)
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "Language updated", result)
}

func (h *languageHandlerImpl) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("invalid language id"))
		return
	}
	if err := h.service.Delete(id); err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "Language deleted", nil)
}
