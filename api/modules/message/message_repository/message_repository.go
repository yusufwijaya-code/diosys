package message_repository

import (
	"portfolio-api/modules/message/message_model"

	"github.com/jmoiron/sqlx"
)

const messageColumns = `messageID, clientName, clientEmail, clientPhone, subject, messageBody, isRead, isArchived, createdDate`

// MessageRepository handles persistence for client messages.
type MessageRepository interface {
	Create(message message_model.ClientMessage) (int, error)
	FindAll() ([]message_model.ClientMessage, error)
	FindByID(messageID int) (message_model.ClientMessage, error)
	UpdateStatus(messageID, isRead, isArchived int) error
	Delete(messageID int) error
}

type messageRepositoryImpl struct {
	db *sqlx.DB
}

// NewMessageRepository builds a MessageRepository.
func NewMessageRepository(db *sqlx.DB) MessageRepository {
	return &messageRepositoryImpl{db: db}
}

func (r *messageRepositoryImpl) Create(message message_model.ClientMessage) (int, error) {
	query := `INSERT INTO ms_client_message (clientName, clientEmail, clientPhone, subject, messageBody) VALUES (?, ?, ?, ?, ?)`
	result, err := r.db.Exec(query, message.ClientName, message.ClientEmail, message.ClientPhone, message.Subject, message.MessageBody)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return int(id), err
}

func (r *messageRepositoryImpl) FindAll() ([]message_model.ClientMessage, error) {
	messages := []message_model.ClientMessage{}
	query := `SELECT ` + messageColumns + ` FROM ms_client_message ORDER BY createdDate DESC, messageID DESC`
	err := r.db.Select(&messages, query)
	return messages, err
}

func (r *messageRepositoryImpl) FindByID(messageID int) (message_model.ClientMessage, error) {
	var message message_model.ClientMessage
	query := `SELECT ` + messageColumns + ` FROM ms_client_message WHERE messageID = ? LIMIT 1`
	err := r.db.Get(&message, query, messageID)
	return message, err
}

func (r *messageRepositoryImpl) UpdateStatus(messageID, isRead, isArchived int) error {
	query := `UPDATE ms_client_message SET isRead = ?, isArchived = ? WHERE messageID = ?`
	_, err := r.db.Exec(query, isRead, isArchived, messageID)
	return err
}

func (r *messageRepositoryImpl) Delete(messageID int) error {
	_, err := r.db.Exec(`DELETE FROM ms_client_message WHERE messageID = ?`, messageID)
	return err
}
