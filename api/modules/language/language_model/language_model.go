package language_model

import "gopkg.in/guregu/null.v4"

const TableName = "ms_language"

// Language represents a spoken language entry of a developer.
type Language struct {
	LanguageID  int       `db:"languageID"  json:"languageID"`
	UserID      int       `db:"userID"      json:"userID"`
	Name        string    `db:"name"        json:"name"`
	Level       string    `db:"level"       json:"level"`
	Icon        string    `db:"icon"        json:"icon"`
	OrderNo     int       `db:"orderNo"     json:"orderNo"`
	CreatedDate null.Time `db:"createdDate" json:"createdDate"`
	EditedDate  null.Time `db:"editedDate"  json:"editedDate"`
}
