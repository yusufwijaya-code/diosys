package message_dto

// MessageRequest is the public contact-form submission payload.
type MessageRequest struct {
	ClientName  string `json:"clientName" binding:"required"`
	ClientEmail string `json:"clientEmail" binding:"required,email"`
	ClientPhone string `json:"clientPhone"`
	Subject     string `json:"subject"`
	MessageBody string `json:"messageBody" binding:"required"`
}

// MessageStatusRequest updates the read/archived flags of a message.
type MessageStatusRequest struct {
	IsRead     *bool `json:"isRead"`
	IsArchived *bool `json:"isArchived"`
}
