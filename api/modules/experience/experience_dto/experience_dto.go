package experience_dto

// ExperienceRequest is the create/update payload for a work experience.
type ExperienceRequest struct {
	Position         string   `json:"position" binding:"required"`
	Company          string   `json:"company"  binding:"required"`
	Period           string   `json:"period"`
	OrderNo          int      `json:"orderNo"`
	Technologies     []string `json:"technologies"`
	Responsibilities []string `json:"responsibilities"`
}

// ExperienceResponse is the aggregate returned to clients.
type ExperienceResponse struct {
	ExperienceID     int      `json:"experienceID"`
	Position         string   `json:"position"`
	Company          string   `json:"company"`
	Period           string   `json:"period"`
	OrderNo          int      `json:"orderNo"`
	Technologies     []string `json:"technologies"`
	Responsibilities []string `json:"responsibilities"`
}
