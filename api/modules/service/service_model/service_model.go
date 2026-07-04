package service_model

import "gopkg.in/guregu/null.v4"

const TableName = "ms_service"

// Service represents an agency capability shown on the landing page.
type Service struct {
	ServiceID   int         `db:"serviceID"   json:"serviceID"`
	Title       string      `db:"title"       json:"title"`
	Description null.String `db:"description" json:"description"`
	Icon        null.String `db:"icon"        json:"icon"`
	OrderNo     int         `db:"orderNo"     json:"orderNo"`
	FlagActive  int         `db:"flagActive"  json:"flagActive"`
	CreatedDate null.Time   `db:"createdDate" json:"createdDate"`
	EditedDate  null.Time   `db:"editedDate"  json:"editedDate"`
}
