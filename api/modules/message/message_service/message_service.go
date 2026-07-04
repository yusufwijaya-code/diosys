package message_service

import (
	"database/sql"
	"errors"

	"portfolio-api/base/helpers/error_helper"
	"portfolio-api/modules/message/message_dto"
	"portfolio-api/modules/message/message_model"
	"portfolio-api/modules/message/message_repository"

	"gopkg.in/guregu/null.v4"
)

// MessageService exposes client message operations.
type MessageService interface {
	Submit(request message_dto.MessageRequest) (message_model.ClientMessage, error)
	GetAll() ([]message_model.ClientMessage, error)
	GetByID(messageID int) (message_model.ClientMessage, error)
	UpdateStatus(messageID int, request message_dto.MessageStatusRequest) (message_model.ClientMessage, error)
	Delete(messageID int) error
}

type messageServiceImpl struct {
	repository message_repository.MessageRepository
}

// NewMessageService builds a MessageService.
func NewMessageService(repository message_repository.MessageRepository) MessageService {
	return &messageServiceImpl{repository: repository}
}

func (s *messageServiceImpl) Submit(request message_dto.MessageRequest) (message_model.ClientMessage, error) {
	id, err := s.repository.Create(message_model.ClientMessage{
		ClientName:  request.ClientName,
		ClientEmail: request.ClientEmail,
		ClientPhone: null.NewString(request.ClientPhone, request.ClientPhone != ""),
		Subject:     null.NewString(request.Subject, request.Subject != ""),
		MessageBody: request.MessageBody,
	})
	if err != nil {
		return message_model.ClientMessage{}, error_helper.Internal(err)
	}
	return s.GetByID(id)
}

func (s *messageServiceImpl) GetAll() ([]message_model.ClientMessage, error) {
	messages, err := s.repository.FindAll()
	if err != nil {
		return nil, error_helper.Internal(err)
	}
	return messages, nil
}

func (s *messageServiceImpl) GetByID(messageID int) (message_model.ClientMessage, error) {
	message, err := s.repository.FindByID(messageID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return message_model.ClientMessage{}, error_helper.NotFound("message not found")
		}
		return message_model.ClientMessage{}, error_helper.Internal(err)
	}
	return message, nil
}

func (s *messageServiceImpl) UpdateStatus(messageID int, request message_dto.MessageStatusRequest) (message_model.ClientMessage, error) {
	existing, err := s.GetByID(messageID)
	if err != nil {
		return message_model.ClientMessage{}, err
	}

	isRead := existing.IsRead
	if request.IsRead != nil {
		isRead = boolToInt(*request.IsRead)
	}
	isArchived := existing.IsArchived
	if request.IsArchived != nil {
		isArchived = boolToInt(*request.IsArchived)
	}

	if err := s.repository.UpdateStatus(messageID, isRead, isArchived); err != nil {
		return message_model.ClientMessage{}, error_helper.Internal(err)
	}
	return s.GetByID(messageID)
}

func (s *messageServiceImpl) Delete(messageID int) error {
	if _, err := s.GetByID(messageID); err != nil {
		return err
	}
	if err := s.repository.Delete(messageID); err != nil {
		return error_helper.Internal(err)
	}
	return nil
}

func boolToInt(value bool) int {
	if value {
		return 1
	}
	return 0
}
