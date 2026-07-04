package message_handler

import (
	"strconv"

	"portfolio-api/base/helpers/error_helper"
	"portfolio-api/base/helpers/http_helper"
	"portfolio-api/constants"
	"portfolio-api/modules/message/message_dto"
	"portfolio-api/modules/message/message_service"

	"github.com/gin-gonic/gin"
)

// MessageHandler exposes contact-form and inbox endpoints.
type MessageHandler interface {
	Submit(c *gin.Context)
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	UpdateStatus(c *gin.Context)
	Delete(c *gin.Context)
}

type messageHandlerImpl struct {
	service message_service.MessageService
}

// NewMessageHandler builds a MessageHandler.
func NewMessageHandler(service message_service.MessageService) MessageHandler {
	return &messageHandlerImpl{service: service}
}

func (h *messageHandlerImpl) Submit(c *gin.Context) {
	var request message_dto.MessageRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("name, a valid email and message are required"))
		return
	}
	result, err := h.service.Submit(request)
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_CREATED, "Message sent", result)
}

func (h *messageHandlerImpl) GetAll(c *gin.Context) {
	result, err := h.service.GetAll()
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "OK", result)
}

func (h *messageHandlerImpl) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("invalid message id"))
		return
	}
	result, err := h.service.GetByID(id)
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "OK", result)
}

func (h *messageHandlerImpl) UpdateStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("invalid message id"))
		return
	}
	var request message_dto.MessageStatusRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("invalid status payload"))
		return
	}
	result, err := h.service.UpdateStatus(id, request)
	if err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "Message updated", result)
}

func (h *messageHandlerImpl) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		http_helper.HttpErrorResponse(c, error_helper.Validation("invalid message id"))
		return
	}
	if err := h.service.Delete(id); err != nil {
		http_helper.HttpErrorResponse(c, err)
		return
	}
	http_helper.SuccessResponse(c, constants.EC_SUCCESS, "Message deleted", nil)
}
