package setting_model

import "gopkg.in/guregu/null.v4"

const TableName = "ms_system_setting"

// Setting is a single global configuration key/value pair.
type Setting struct {
	SettingID    int         `db:"settingID"    json:"settingID"`
	SettingKey   string      `db:"settingKey"   json:"settingKey"`
	SettingValue null.String `db:"settingValue" json:"settingValue"`
	Description  null.String `db:"description"  json:"description"`
	EditedDate   null.Time   `db:"editedDate"   json:"editedDate"`
}
