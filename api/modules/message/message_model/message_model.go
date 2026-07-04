package message_model

import "gopkg.in/guregu/null.v4"

const TableName = "ms_client_message"

// ClientMessage represents an inbound message from the public contact form.
type ClientMessage struct {
	MessageID   int         `db:"messageID"   json:"messageID"`
	ClientName  string      `db:"clientName"  json:"clientName"`
	ClientEmail string      `db:"clientEmail" json:"clientEmail"`
	ClientPhone null.String `db:"clientPhone" json:"clientPhone"`
	Subject     null.String `db:"subject"     json:"subject"`
	MessageBody string      `db:"messageBody" json:"messageBody"`
	IsRead      int         `db:"isRead"      json:"isRead"`
	IsArchived  int         `db:"isArchived"  json:"isArchived"`
	CreatedDate null.Time   `db:"createdDate" json:"createdDate"`
}
