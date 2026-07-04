package language_dto

// LanguageRequest is the create/update payload for a language.
type LanguageRequest struct {
	Name    string `json:"name"    binding:"required"`
	Level   string `json:"level"   binding:"required"`
	Icon    string `json:"icon"`
	OrderNo int    `json:"orderNo"`
}
