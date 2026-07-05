package testimonial_handler

import (
	"strconv"

	"portfolio-api/base/helpers/error_helper"
	"portfolio-api/base/helpers/http_helper"
	"portfolio-api/constants"
	"portfolio-api/modules/testimonial/testimonial_dto"
	"portfolio-api/modules/testimonial/testimonial_service"

	"github.com/gin-gonic/gin"
)

type TestimonialHandler interface {
	GetPublic(c *gin.Context)
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	UploadPhoto(c *gin.Context)
	Delete(c *gin.Context)
}

type testimonialHandlerImpl struct{ service testimonial_service.TestimonialService }

func NewTestimonialHandler(service testimonial_service.TestimonialService) TestimonialHandler {
	return &testimonialHandlerImpl{service: service}
}

func (h *testimonialHandlerImpl) GetPublic(c *gin.Context) {
	result, err := h.service.GetPublic()
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "OK", result)
}

func (h *testimonialHandlerImpl) GetAll(c *gin.Context) {
	result, err := h.service.GetAll()
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "OK", result)
}

func (h *testimonialHandlerImpl) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("invalid id"))
		return
	}
	result, err := h.service.GetByID(id)
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "OK", result)
}

func (h *testimonialHandlerImpl) Create(c *gin.Context) {
	var req testimonial_dto.TestimonialRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("clientName and testimonialText are required"))
		return
	}
	result, err := h.service.Create(req)
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_CREATED, "Testimonial created", result)
}

func (h *testimonialHandlerImpl) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("invalid id"))
		return
	}
	var req testimonial_dto.TestimonialRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("clientName and testimonialText are required"))
		return
	}
	result, err := h.service.Update(id, req)
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "Testimonial updated", result)
}

func (h *testimonialHandlerImpl) UploadPhoto(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("invalid id"))
		return
	}
	fileHeader, err := c.FormFile("file")
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("file is required"))
		return
	}
	result, err := h.service.UploadPhoto(id, fileHeader)
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "Photo uploaded", result)
}

func (h *testimonialHandlerImpl) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("invalid id"))
		return
	}
	if err := h.service.Delete(id); err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "Testimonial deleted", nil)
}
