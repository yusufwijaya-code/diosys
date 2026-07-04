package message

import (
	"portfolio-api/modules/message/message_handler"
	"portfolio-api/modules/message/message_repository"
	"portfolio-api/modules/message/message_service"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// Router wires and registers the contact-form and inbox endpoints.
func Router(public *gin.RouterGroup, protected *gin.RouterGroup, db *sqlx.DB) {
	repository := message_repository.NewMessageRepository(db)
	service := message_service.NewMessageService(repository)
	handler := message_handler.NewMessageHandler(service)

	public.POST("/contact", handler.Submit)

	protectedGroup := protected.Group("/messages")
	{
		protectedGroup.GET("", handler.GetAll)
		protectedGroup.GET("/:id", handler.GetByID)
		protectedGroup.PATCH("/:id/status", handler.UpdateStatus)
		protectedGroup.DELETE("/:id", handler.Delete)
	}
}
