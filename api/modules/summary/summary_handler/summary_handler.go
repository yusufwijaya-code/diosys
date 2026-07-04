package summary_handler

import (
	"strconv"

	"portfolio-api/base/helpers/error_helper"
	"portfolio-api/base/helpers/http_helper"
	"portfolio-api/constants"
	"portfolio-api/modules/summary/summary_dto"
	"portfolio-api/modules/summary/summary_service"

	"github.com/gin-gonic/gin"
)

// SummaryHandler exposes summary endpoints scoped to a developer.
type SummaryHandler interface {
	Get(c *gin.Context)
	Save(c *gin.Context)
}

type summaryHandlerImpl struct {
	service summary_service.SummaryService
}

// NewSummaryHandler builds a SummaryHandler.
func NewSummaryHandler(service summary_service.SummaryService) SummaryHandler {
	return &summaryHandlerImpl{service: service}
}

func (h *summaryHandlerImpl) Get(c *gin.Context) {
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

func (h *summaryHandlerImpl) Save(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("invalid developer id"))
		return
	}
	var request summary_dto.SummaryRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("content is required"))
		return
	}
	result, err := h.service.Save(userID, request)
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "Summary saved", result)
}
