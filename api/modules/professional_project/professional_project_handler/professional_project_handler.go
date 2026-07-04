package professional_project_handler

import (
	"strconv"

	"portfolio-api/base/helpers/error_helper"
	"portfolio-api/base/helpers/http_helper"
	"portfolio-api/constants"
	"portfolio-api/modules/professional_project/professional_project_dto"
	"portfolio-api/modules/professional_project/professional_project_service"

	"github.com/gin-gonic/gin"
)

type ProfessionalProjectHandler interface {
	GetPublicByID(c *gin.Context)
	GetByUser(c *gin.Context)
	Create(c *gin.Context)
	Delete(c *gin.Context)
	UploadThumbnail(c *gin.Context)
	AddFeature(c *gin.Context)
	DeleteFeature(c *gin.Context)
	AddFeatureImage(c *gin.Context)
	DeleteFeatureImage(c *gin.Context)
}

type profProjHandlerImpl struct {
	service professional_project_service.ProfessionalProjectService
}

func NewProfessionalProjectHandler(service professional_project_service.ProfessionalProjectService) ProfessionalProjectHandler {
	return &profProjHandlerImpl{service: service}
}

func (h *profProjHandlerImpl) GetPublicByID(c *gin.Context) {
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

func (h *profProjHandlerImpl) GetByUser(c *gin.Context) {
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

func (h *profProjHandlerImpl) Create(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("invalid developer id"))
		return
	}
	var req professional_project_dto.ProfessionalProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("title and company are required"))
		return
	}
	result, err := h.service.Create(userID, req)
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_CREATED, "Professional project created", result)
}

func (h *profProjHandlerImpl) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("invalid project id"))
		return
	}
	if err := h.service.Delete(id); err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "Professional project deleted", nil)
}

func (h *profProjHandlerImpl) UploadThumbnail(c *gin.Context) {
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

func (h *profProjHandlerImpl) AddFeature(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("invalid project id"))
		return
	}
	var req professional_project_dto.ProjectFeatureRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("title is required"))
		return
	}
	result, err := h.service.AddFeature(id, req)
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_CREATED, "Feature added", result)
}

func (h *profProjHandlerImpl) DeleteFeature(c *gin.Context) {
	featureID, err := strconv.Atoi(c.Param("featureID"))
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("invalid feature id"))
		return
	}
	if err := h.service.DeleteFeature(featureID); err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "Feature deleted", nil)
}

func (h *profProjHandlerImpl) AddFeatureImage(c *gin.Context) {
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
	caption := c.PostForm("caption")
	result, err := h.service.AddFeatureImage(featureID, fileHeader, caption)
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_CREATED, "Image added", result)
}

func (h *profProjHandlerImpl) DeleteFeatureImage(c *gin.Context) {
	imageID, err := strconv.Atoi(c.Param("imageID"))
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("invalid image id"))
		return
	}
	if err := h.service.DeleteFeatureImage(imageID); err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "Image deleted", nil)
}
