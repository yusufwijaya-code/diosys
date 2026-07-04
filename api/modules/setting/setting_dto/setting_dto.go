package setting_dto

// SettingItem is a single key/value pair in an update payload.
type SettingItem struct {
	SettingKey   string `json:"settingKey" binding:"required"`
	SettingValue string `json:"settingValue"`
}

// SettingUpdateRequest is the bulk upsert payload for system settings.
type SettingUpdateRequest struct {
	Settings []SettingItem `json:"settings" binding:"required"`
}
