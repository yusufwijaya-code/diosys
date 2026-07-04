package certificate_model

import "gopkg.in/guregu/null.v4"

const TableName = "ms_certificate"

// Certificate represents a developer's certification entry.
type Certificate struct {
	CertificateID int         `db:"certificateID" json:"certificateID"`
	UserID        int         `db:"userID"        json:"userID"`
	Name          string      `db:"name"          json:"name"`
	Issuer        string      `db:"issuer"        json:"issuer"`
	Period        string      `db:"period"        json:"period"`
	Link          null.String `db:"link"          json:"link"`
	OrderNo       int         `db:"orderNo"       json:"orderNo"`
	CreatedDate   null.Time   `db:"createdDate"   json:"createdDate"`
	EditedDate    null.Time   `db:"editedDate"    json:"editedDate"`
}
