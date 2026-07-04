package developer_handler

import (
	"strconv"

	"portfolio-api/base/helpers/error_helper"
	"portfolio-api/base/helpers/http_helper"
	"portfolio-api/constants"
	"portfolio-api/modules/developer/developer_dto"
	"portfolio-api/modules/developer/developer_service"

	"github.com/gin-gonic/gin"
)

// DeveloperHandler exposes developer directory, profile and CMS endpoints.
type DeveloperHandler interface {
	GetDirectory(c *gin.Context)
	GetProfile(c *gin.Context)
	List(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	UploadPhoto(c *gin.Context)
	UploadCV(c *gin.Context)
}

type developerHandlerImpl struct {
	service developer_service.DeveloperService
}

// NewDeveloperHandler builds a DeveloperHandler.
func NewDeveloperHandler(service developer_service.DeveloperService) DeveloperHandler {
	return &developerHandlerImpl{service: service}
}

func (h *developerHandlerImpl) GetDirectory(c *gin.Context) {
	result, err := h.service.GetDirectory()
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "OK", result)
}

func (h *developerHandlerImpl) GetProfile(c *gin.Context) {
	username := c.Param("username")
	if username == "" {
		http_helper.HttpErrorResponse(c, error_helper.Validation("username is required"))
		return
	}
	result, err := h.service.GetProfileByUsername(username)
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "OK", result)
}

func (h *developerHandlerImpl) List(c *gin.Context) {
	result, err := h.service.List()
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "OK", result)
}

func (h *developerHandlerImpl) GetByID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("invalid developer id"))
		return
	}
	result, err := h.service.GetByID(userID)
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "OK", result)
}

func (h *developerHandlerImpl) Create(c *gin.Context) {
	var request developer_dto.DeveloperRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("username, email and full name are required"))
		return
	}
	result, err := h.service.Create(request)
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_CREATED, "Developer created", result)
}

func (h *developerHandlerImpl) Update(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("invalid developer id"))
		return
	}
	var request developer_dto.DeveloperRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("username, email and full name are required"))
		return
	}
	result, err := h.service.Update(userID, request)
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "Developer updated", result)
}

func (h *developerHandlerImpl) Delete(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("invalid developer id"))
		return
	}
	if err := h.service.Delete(userID); err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "Developer deleted", nil)
}

func (h *developerHandlerImpl) UploadPhoto(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("invalid developer id"))
		return
	}
	fileHeader, err := c.FormFile("file")
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("file is required"))
		return
	}
	result, err := h.service.UploadPhoto(userID, fileHeader)
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "Photo uploaded", result)
}

func (h *developerHandlerImpl) UploadCV(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("invalid developer id"))
		return
	}
	fileHeader, err := c.FormFile("file")
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("file is required"))
		return
	}
	result, err := h.service.UploadCV(userID, fileHeader)
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "CV uploaded", result)
}
