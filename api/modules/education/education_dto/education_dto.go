package education_dto

// EducationRequest is the create/update payload for an education entry.
type EducationRequest struct {
	Degree       string   `json:"degree"      binding:"required"`
	Institution  string   `json:"institution" binding:"required"`
	Year         string   `json:"year"`
	Type         string   `json:"type"`
	OrderNo      int      `json:"orderNo"`
	Achievements []string `json:"achievements"`
}

// EducationResponse is the aggregate returned to clients.
type EducationResponse struct {
	EducationID  int      `json:"educationID"`
	Degree       string   `json:"degree"`
	Institution  string   `json:"institution"`
	Year         string   `json:"year"`
	Type         string   `json:"type"`
	OrderNo      int      `json:"orderNo"`
	Achievements []string `json:"achievements"`
}
