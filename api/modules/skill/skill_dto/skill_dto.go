package skill_dto

// SkillRequest is the create/update payload for a skill.
type SkillRequest struct {
	Name     string `json:"name"     binding:"required"`
	Level    string `json:"level"    binding:"required"`
	Category string `json:"category" binding:"required"`
	OrderNo  int    `json:"orderNo"`
}
