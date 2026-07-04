package project_handler

import (
	"strconv"

	"portfolio-api/base/helpers/error_helper"
	"portfolio-api/base/helpers/http_helper"
	"portfolio-api/constants"
	"portfolio-api/modules/project/project_dto"
	"portfolio-api/modules/project/project_service"

	"github.com/gin-gonic/gin"
)

// ProjectHandler exposes project endpoints.
type ProjectHandler interface {
	GetAllPublic(c *gin.Context)
	GetPublicByID(c *gin.Context)
	GetByUser(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	UploadThumbnail(c *gin.Context)
	AddImage(c *gin.Context)
	DeleteImage(c *gin.Context)
	AddFeatureImage(c *gin.Context)
	DeleteFeatureImage(c *gin.Context)
}

type projectHandlerImpl struct {
	service project_service.ProjectService
}

// NewProjectHandler builds a ProjectHandler.
func NewProjectHandler(service project_service.ProjectService) ProjectHandler {
	return &projectHandlerImpl{service: service}
}

func (h *projectHandlerImpl) GetAllPublic(c *gin.Context) {
	result, err := h.service.GetAllPublic()
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "OK", result)
}

func (h *projectHandlerImpl) GetPublicByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("invalid project id"))
		return
	}
	result, err := h.service.GetByID(id)
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "OK", result)
}

func (h *projectHandlerImpl) GetByUser(c *gin.Context) {
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

func (h *projectHandlerImpl) Create(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("invalid developer id"))
		return
	}
	var request project_dto.ProjectRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("title is required"))
		return
	}
	result, err := h.service.Create(userID, request)
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_CREATED, "Project created", result)
}

func (h *projectHandlerImpl) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("invalid project id"))
		return
	}
	var request project_dto.ProjectRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("title is required"))
		return
	}
	result, err := h.service.Update(id, request)
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "Project updated", result)
}

func (h *projectHandlerImpl) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("invalid project id"))
		return
	}
	if err := h.service.Delete(id); err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "Project deleted", nil)
}

func (h *projectHandlerImpl) UploadThumbnail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("invalid project id"))
		return
	}
	fileHeader, err := c.FormFile("file")
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("file is required"))
		return
	}
	result, err := h.service.UploadThumbnail(id, fileHeader)
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "Thumbnail uploaded", result)
}

func (h *projectHandlerImpl) AddImage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("invalid project id"))
		return
	}
	fileHeader, err := c.FormFile("file")
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("file is required"))
		return
	}
	result, err := h.service.AddImage(id, fileHeader, c.PostForm("caption"))
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "Image added", result)
}

func (h *projectHandlerImpl) DeleteImage(c *gin.Context) {
	imageID, err := strconv.Atoi(c.Param("imageID"))
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("invalid image id"))
		return
	}
	if err := h.service.DeleteImage(imageID); err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "Image deleted", nil)
}

func (h *projectHandlerImpl) AddFeatureImage(c *gin.Context) {
	featureID, err := strconv.Atoi(c.Param("featureID"))
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("invalid feature id"))
		return
	}
	fileHeader, err := c.FormFile("file")
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("file is required"))
		return
	}
	result, err := h.service.AddFeatureImage(featureID, fileHeader, c.PostForm("caption"))
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_CREATED, "Feature image added", result)
}

func (h *projectHandlerImpl) DeleteFeatureImage(c *gin.Context) {
	imageID, err := strconv.Atoi(c.Param("imageID"))
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("invalid image id"))
		return
	}
	if err := h.service.DeleteFeatureImage(imageID); err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "Feature image deleted", nil)
}
