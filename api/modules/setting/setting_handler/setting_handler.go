package setting_handler

import (
	"portfolio-api/base/helpers/error_helper"
	"portfolio-api/base/helpers/http_helper"
	"portfolio-api/constants"
	"portfolio-api/modules/setting/setting_dto"
	"portfolio-api/modules/setting/setting_service"

	"github.com/gin-gonic/gin"
)

// SettingHandler exposes system setting endpoints.
type SettingHandler interface {
	GetPublic(c *gin.Context)
	GetAll(c *gin.Context)
	Update(c *gin.Context)
}

type settingHandlerImpl struct {
	service setting_service.SettingService
}

// NewSettingHandler builds a SettingHandler.
func NewSettingHandler(service setting_service.SettingService) SettingHandler {
	return &settingHandlerImpl{service: service}
}

func (h *settingHandlerImpl) GetPublic(c *gin.Context) {
	result, err := h.service.GetPublicMap()
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "OK", result)
}

func (h *settingHandlerImpl) GetAll(c *gin.Context) {
	result, err := h.service.GetAll()
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "OK", result)
}

func (h *settingHandlerImpl) Update(c *gin.Context) {
	var request setting_dto.SettingUpdateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("settings payload is required"))
		return
	}
	result, err := h.service.Update(request)
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "Settings updated", result)
}
